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

	"main/my_redis"
)

func main() {
	myRedis := my_redis.NewMyRedis(2)

	myRedis.Add("a", 1)
	myRedis.Add("b", 2)
	myRedis.Add("c", 3) // This should evict "a"

	if _, ok := myRedis.Get("a"); ok {
		fmt.Printf("expected key 'a' to be evicted\n")
	}

	if value, ok := myRedis.Get("b"); !ok || value != 2 {
		fmt.Printf("expected key 'b' to have value 2, got %d\n", value)
	}

	if value, ok := myRedis.Get("c"); !ok || value != 3 {
		fmt.Printf("expected key 'c' to have value 3, got %d\n", value)
	}
}

func test_concurrency(my_redis *my_redis.MyRedis) {
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
