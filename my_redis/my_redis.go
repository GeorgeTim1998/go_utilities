package my_redis

import (
	"sync"
	"time"
)

type CacheItem struct {
	value      interface{}
	expiration *time.Time
}

type MyRedis struct {
	mu    sync.Mutex
	cache map[interface{}]CacheItem
	cap   int
	order []interface{}
}

func NewMyRedis(cap int) *MyRedis {
	return &MyRedis{
		cache: make(map[interface{}]CacheItem, cap),
		cap:   cap,
		order: make([]interface{}, 0, cap),
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
	my_redis.cache = make(map[interface{}]CacheItem, my_redis.cap)
	my_redis.order = make([]interface{}, 0, my_redis.cap)
}

func (my_redis *MyRedis) Add(key, value interface{}) {
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

func (my_redis *MyRedis) Get(key interface{}) (interface{}, bool) {
	my_redis.mu.Lock()
	defer my_redis.mu.Unlock()

	item, exists := my_redis.cache[key]
	if !exists || (item.expiration != nil && time.Now().After(*item.expiration)) {
		return nil, false
	}

	my_redis.removeFromOrder(key)
	my_redis.order = append(my_redis.order, key)

	return item.value, true
}

func (my_redis *MyRedis) Remove(key interface{}) {
	my_redis.mu.Lock()
	defer my_redis.mu.Unlock()

	delete(my_redis.cache, key)
	my_redis.removeFromOrder(key)
}

func (my_redis *MyRedis) AddWithTTL(key, value interface{}, ttl time.Duration) {
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

func (my_redis *MyRedis) removeFromOrder(key interface{}) {
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

	delete(my_redis.cache, lruKey)
	my_redis.removeFromOrder(lruKey)
}
