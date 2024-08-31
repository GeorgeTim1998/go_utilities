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

func (my_redis *MyRedis) Cap() int {
	return my_redis.cap
}

func (my_redis *MyRedis) Len() int {
	return len(my_redis.cache)
}

func (my_redis *MyRedis) Clear() {
	my_redis.cache = make(map[string]int, 10)
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
	my_redis.mu.Lock()
	defer my_redis.mu.Unlock()

	my_redis.Add(key, value)
	my_redis.deleteAfterTTL(key, ttl)
}

func (my_redis *MyRedis) deleteAfterTTL(key string, ttl time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), ttl)

	go func() {
		defer cancel()

		for range ctx.Done() {
			my_redis.Remove(key)
			return
		}
	}()
}

func main() {
	my_redis := MyRedis{mu: sync.Mutex{}, cache: make(map[string]int), cap: 4}
	my_redis.cache["a"] = 1

	fmt.Println(my_redis.Len())
}
