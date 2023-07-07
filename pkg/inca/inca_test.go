package inca

import (
	"strconv"
	"testing"
)

// Benchmarking inca
func BenchmarkSetSameKey(b *testing.B) {

	cache := CreateInca(&IncaConfig{
		Capacity:       10000,
		EvictionPolicy: TYPE_LRU,
	})

	for i := 0; i <= b.N; i++ {
		cache.Memory.Set("key", strconv.FormatInt(int64(i), 10), -1)
	}
}

func BenchmarkSetKeyTillCapacity(b *testing.B) {
	// Not good results here on setting CAPACITY to 100000
	const CAPACITY = 10000
	cache := CreateInca(&IncaConfig{
		Capacity:       CAPACITY,
		EvictionPolicy: TYPE_LRU,
	})

	for i := 0; i <= CAPACITY; i++ {
		kv := strconv.FormatInt(int64(i), 10)
		cache.Memory.Set(kv, kv, -1)
	}
}

func BenchmarkSetKeyfor1MoreThanCapacity(b *testing.B) {
	const CAPACITY = 1
	cache := CreateInca(&IncaConfig{
		Capacity:       CAPACITY,
		EvictionPolicy: TYPE_LRU,
	})

	for i := 0; i <= CAPACITY+1; i++ {
		kv := strconv.FormatInt(int64(i), 10)
		cache.Memory.Set(kv, kv, -1)
	}
}

func BenchmarkSetKeyfor100MoreThanCapacity(b *testing.B) {
	const CAPACITY = 1
	cache := CreateInca(&IncaConfig{
		Capacity:       CAPACITY,
		EvictionPolicy: TYPE_LRU,
	})

	for i := 0; i <= CAPACITY+100; i++ {
		kv := strconv.FormatInt(int64(i), 10)
		cache.Memory.Set(kv, kv, -1)
	}
}

func BenchmarkSetKeyfor10xCapacity(b *testing.B) {
	const CAPACITY = 10000
	cache := CreateInca(&IncaConfig{
		Capacity:       CAPACITY,
		EvictionPolicy: TYPE_LRU,
	})

	for i := 0; i <= CAPACITY*10; i++ {
		kv := strconv.FormatInt(int64(i), 10)
		cache.Memory.Set(kv, kv, -1)
	}
}
func BenchmarkSetKey(b *testing.B) {
	const CAPACITY = 10

	cache := CreateInca(&IncaConfig{
		Capacity:       CAPACITY,
		EvictionPolicy: TYPE_LRU,
	})

	cache.Memory.Set("key", "value", -1)
}

func BenchmarkSetGetKey(b *testing.B) {
	const CAPACITY = 10

	cache := CreateInca(&IncaConfig{
		Capacity:       CAPACITY,
		EvictionPolicy: TYPE_LRU,
	})

	cache.Memory.Set("key", "value", -1)
	cache.Memory.Get("key")
}
