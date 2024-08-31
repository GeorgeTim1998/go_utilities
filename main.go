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
	"context"
	"fmt"
	"sync"
	"time"
)

type MyRedis struct {
	mu    sync.Mutex
	cache map[string]int
	cap   int
}

func NewMyRedis(cap int) *MyRedis {
	return &MyRedis{
		cache: make(map[string]int, cap),
		cap:   cap,
	}
}

func (my_redis *MyRedis) Cap() int {
	return my_redis.cap
}

func (my_redis *MyRedis) Len() int {
	my_redis.mu.Lock()
	defer my_redis.mu.Unlock()

	return len(my_redis.cache)
}

func (my_redis *MyRedis) Clear() {
	my_redis.mu.Lock()
	defer my_redis.mu.Unlock()

	my_redis.cache = make(map[string]int, my_redis.cap) // what is faster: this or one by one?
}

func (my_redis *MyRedis) Add(key string, value int) {
	my_redis.mu.Lock()
	defer my_redis.mu.Unlock()

	my_redis.cache[key] = value
}

func (my_redis *MyRedis) Get(key string) (int, bool) {
	my_redis.mu.Lock()
	defer my_redis.mu.Unlock()

	value, ok := my_redis.cache[key]

	return value, ok
}

func (my_redis *MyRedis) Remove(key string) {
	my_redis.mu.Lock()
	defer my_redis.mu.Unlock()

	delete(my_redis.cache, key)
}

func (my_redis *MyRedis) AddWithTTL(key string, value int, ttl time.Duration) {
	my_redis.Add(key, value)
	my_redis.deleteAfterTTL(key, ttl)
}

func (my_redis *MyRedis) deleteAfterTTL(key string, ttl time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), ttl)

	go func() {
		defer cancel()

		<-ctx.Done()
		my_redis.Remove(key)
	}()
}

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
