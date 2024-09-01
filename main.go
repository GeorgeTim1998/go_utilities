// я пришлю задание, ты сможешь оценить срок выполнения

// Задание: ""Написать сервис (класс/структура) кэширования.

// Основные условия:
// - Кэш ограниченной емкости, метод вытеснения ключей LRU;
// - Сервис должен быть потокобезопасный;
// - Сервис должен принимать любые значения;
// - Реализовать unit-тесты.

// Сервис должен реализовывать следующий интерфейс:
//     type ICache interface {
//         Cap() int
//         Len() int
//         Clear() // удаляет все ключи
//         Add(key, value any)
//         AddWithTTL(key, value any, ttl time.Duration) // добавляет ключ со сроком жизни ttl
//         Get(key any) (value any, ok bool)
//         Remove(key any)
//     }

// TTL - через сколько должен удалиться ключ
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	my_redis := NewMyRedis(4)

	my_redis.AddWithTTL("a", 1, 5*time.Second)

	for i := 0; i < 1000; i++ {
		value, ok := my_redis.Get("a")
		fmt.Println(value, ok)
		time.Sleep(time.Second)
	}
}

func test_concurrency(my_redis *MyRedis) {
	wg := sync.WaitGroup{}
	attempts := 1000

	wg.Add(2 * attempts)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			fmt.Print("a")
			my_redis.Add("a", 1)
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			fmt.Print("b")
			my_redis.Remove("a")
		}()
	}

	wg.Wait()
}
