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

type CacheItem struct {
	value      int
	expiration *time.Time
}

type MyRedis struct {
	mu    sync.Mutex
	cache map[string]CacheItem
	cap   int
	order []string
}

func NewMyRedis(cap int) *MyRedis {
	return &MyRedis{
		cache: make(map[string]CacheItem, cap),
		cap:   cap,
		order: make([]string, 0, cap),
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
	my_redis.cache = make(map[string]CacheItem, my_redis.cap) // What is faster: make new or delete by key
	my_redis.order = make([]string, 0, my_redis.cap)
}

func (my_redis *MyRedis) Add(key string, value int) {
	my_redis.mu.Lock()
	defer my_redis.mu.Unlock()

	if _, exists := my_redis.cache[key]; exists {
		my_redis.removeFromOrder(key)
	}

	my_redis.cache[key] = CacheItem{value: value, expiration: nil}
	my_redis.order = append(my_redis.order, key)

	if len(my_redis.cache) > my_redis.cap {
		my_redis.evictLRU()
	}
}

func (my_redis *MyRedis) Get(key string) (int, bool) {
	my_redis.mu.Lock()
	defer my_redis.mu.Unlock()

	item, exists := my_redis.cache[key]
	if !exists || (item.expiration != nil && time.Now().After(*item.expiration)) {
		return 0, false
	}

	my_redis.removeFromOrder(key)
	my_redis.order = append(my_redis.order, key)

	return item.value, true
}

func (my_redis *MyRedis) Remove(key string) {
	my_redis.mu.Lock()
	defer my_redis.mu.Unlock()

	delete(my_redis.cache, key)
	my_redis.removeFromOrder(key)
}

func (my_redis *MyRedis) AddWithTTL(key string, value int, ttl time.Duration) {
	my_redis.mu.Lock()
	defer my_redis.mu.Unlock()

	expiration := time.Now().Add(ttl)

	if _, exists := my_redis.cache[key]; exists {
		my_redis.removeFromOrder(key)
	}

	my_redis.cache[key] = CacheItem{value: value, expiration: &expiration}
	my_redis.order = append(my_redis.order, key)

	if len(my_redis.cache) > my_redis.cap {
		my_redis.evictLRU()
	}
}

func (my_redis *MyRedis) removeFromOrder(key string) {
	for i, k := range my_redis.order {
		if k == key {
			my_redis.order = append(my_redis.order[:i], my_redis.order[i+1:]...)
			break
		}
	}
}

func (my_redis *MyRedis) evictLRU() {
	if len(my_redis.order) == 0 {
		return
	}

	lruKey := my_redis.order[0]
	my_redis.Remove(lruKey)
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
