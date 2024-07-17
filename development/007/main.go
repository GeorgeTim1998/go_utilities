package main

import (
	"fmt"
	"time"
)

// Функция or объединяет один или более done-каналов в один single-канал
var or func(channels ...<-chan interface{}) <-chan interface{}

func init() {
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			// Если нет каналов, возвращаем закрытый канал
			closedChan := make(chan interface{})
			close(closedChan)
			return closedChan
		case 1:
			// Если один канал, возвращаем его
			return channels[0]
		}

		// Создаем канал, который будем возвращать
		orDone := make(chan interface{})
		go func() {
			defer close(orDone)

			// Используем select для отслеживания закрытия любого из каналов
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				// вызываем рекурсивно, чтобы следить за всеми каналами
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}
}

func main() {
	// Функция, создающая канал, который закрывается через заданное время
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	// Используем функцию or для объединения нескольких каналов
	<-or(
		sig(1*time.Second),
		sig(5*time.Minute),
		sig(2*time.Hour),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
