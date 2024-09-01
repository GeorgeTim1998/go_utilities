package my_redis

import (
	"testing"
	"time"
)

func TestNewMyRedis(t *testing.T) {
	cap := 4
	myRedis := NewMyRedis(cap)

	if myRedis.Cap() != cap {
		t.Errorf("expected cap %d, got %d", cap, myRedis.Cap())
	}

	if myRedis.Len() != 0 {
		t.Errorf("expected length 0, got %d", myRedis.Len())
	}
}

func TestAddAndGet(t *testing.T) {
	myRedis := NewMyRedis(4)
	myRedis.Add("a", 1)

	value, ok := myRedis.Get("a")
	if !ok {
		t.Errorf("expected key 'a' to be present")
	}
	if value != 1 {
		t.Errorf("expected value 1, got %d", value)
	}
}

func TestEviction(t *testing.T) {
	myRedis := NewMyRedis(2)

	myRedis.Add("a", 1)
	myRedis.Add("b", 2)
	myRedis.Add("c", 3) // This should evict "a"

	if _, ok := myRedis.Get("a"); ok {
		t.Errorf("expected key 'a' to be evicted")
	}

	if value, ok := myRedis.Get("b"); !ok || value != 2 {
		t.Errorf("expected key 'b' to have value 2, got %d", value)
	}

	if value, ok := myRedis.Get("c"); !ok || value != 3 {
		t.Errorf("expected key 'c' to have value 3, got %d", value)
	}
}

func TestAddWithTTL(t *testing.T) {
	myRedis := NewMyRedis(4)
	myRedis.AddWithTTL("a", 1, 1*time.Second)

	time.Sleep(2 * time.Second)
	if _, ok := myRedis.Get("a"); ok {
		t.Errorf("expected key 'a' to be expired and removed")
	}
}

func TestRemove(t *testing.T) {
	myRedis := NewMyRedis(4)
	myRedis.Add("a", 1)
	myRedis.Remove("a")

	if _, ok := myRedis.Get("a"); ok {
		t.Errorf("expected key 'a' to be removed")
	}
}

func TestClear(t *testing.T) {
	myRedis := NewMyRedis(4)
	myRedis.Add("a", 1)
	myRedis.Add("b", 2)
	myRedis.Clear()

	if myRedis.Len() != 0 {
		t.Errorf("expected length 0, got %d", myRedis.Len())
	}

	myRedis.Add("c", 3)
	if myRedis.Len() != 1 {
		t.Errorf("expected length 1, got %d", myRedis.Len())
	}
}
