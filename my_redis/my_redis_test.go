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

	_, ok = myRedis.Get("b")
	if ok {
		t.Errorf("expected key 'b' to be abscent")
	}
}

func TestAnyKeyAndValueAddAndGet(t *testing.T) {
	myRedis := NewMyRedis(5)

	tests := []struct {
		key   interface{}
		value interface{}
	}{
		{key: "string_key", value: "string_value"},
		{key: 42, value: 100},
		{key: 3.14, value: "pi"},
		{key: true, value: false},
		{key: struct{ Name string }{Name: "test"}, value: [3]int{1, 2, 3}},
	}

	for _, tt := range tests {
		myRedis.Add(tt.key, tt.value)

		if val, ok := myRedis.Get(tt.key); !ok || val != tt.value {
			t.Errorf("Failed to get the correct value for key: %v, expected: %v, got: %v", tt.key, tt.value, val)
		}
	}
}

func TestAddWithDifferentTTL(t *testing.T) {
	myRedis := NewMyRedis(2)
	myRedis.AddWithTTL("a", 1, time.Second)

	myRedis.AddWithTTL("a", 1, time.Microsecond)
	time.Sleep(2 * time.Microsecond)
	if _, ok := myRedis.Get("a"); ok {
		t.Errorf("expected key 'a' to be expired and removed")
	}
}

func TestAddWithNoTTL(t *testing.T) {
	myRedis := NewMyRedis(2)
	myRedis.AddWithTTL("a", 1, 500*time.Microsecond)
	myRedis.Add("a", 1)

	time.Sleep(700 * time.Microsecond)
	if _, ok := myRedis.Get("a"); !ok {
		t.Errorf("expected key 'a' to be present")
	}
}

func TestAnyKeyAndValueAddWithTTL(t *testing.T) {
	myRedis := NewMyRedis(5)

	tests := []struct {
		key   interface{}
		value interface{}
	}{
		{key: "string_key", value: "string_value"},
		{key: 42, value: 100},
		{key: 3.14, value: "pi"},
		{key: true, value: false},
		{key: struct{ Name string }{Name: "test"}, value: [3]int{1, 2, 3}},
	}

	for _, tt := range tests {
		myRedis.AddWithTTL(tt.key, tt.value, 500*time.Microsecond)

		if val, ok := myRedis.Get(tt.key); !ok || val != tt.value {
			t.Errorf("Failed to get the correct value for key: %v, expected: %v, got: %v", tt.key, tt.value, val)
		}

		time.Sleep(700 * time.Microsecond)

		if _, ok := myRedis.Get(tt.key); ok {
			t.Errorf("expected key %v to be expired and removed", tt.key)
		}
	}
}

func TestEviction(t *testing.T) {
	myRedis := NewMyRedis(2)

	myRedis.Add("a", 1)
	myRedis.Add("b", 2)
	myRedis.Add("c", 3)

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
