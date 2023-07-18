package lru

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Set tests
func TestLRUCacheSet(t *testing.T) {
	/*
		Objectives:

		- Create a cache with a specific capacity 'C'
		- Set key:value as numbers from 1 to C+1 as string
		- Now check whether
			- the last element in the keys is "2"
			- the first element in the keys is "C"
	*/
	const C = 10
	cache := CreateLRUCache(C)

	for i := 1; i <= C+1; i++ {
		kv := strconv.FormatInt(int64(i), 10)
		cache.Set(kv, kv, -1)
	}

	if cache.keys.Tail().Val != strconv.FormatInt(int64(2), 10) {
		t.Errorf("LRUCacheSet test failed! output: %s, expected: %s.\n", cache.keys.Tail().Val, "2")
	}

	if cache.keys.Head().Val != strconv.FormatInt(int64(C+1), 10) {
		t.Errorf("LRUCacheSet test failed! output: %s, expected: %s.\n", cache.keys.Head().Val, strconv.FormatInt(int64(C+1), 10))
	}
}

func TestLRUSetLengthNoTTL(t *testing.T) {
	const C = 3
	cache := CreateLRUCache(C)

	cache.Set("name", "rajab", -1)
	cache.Set("age", "17", -1)
	cache.Set("job", "swe", -1)
	t.Logf("Length: %d\n", cache.keys.LengthC)
	assert.Equal(t, uint32(3), cache.keys.LengthC)
}

func TestLRUSetLengthTTL(t *testing.T) {
	const C = 3
	cache := CreateLRUCache(C)

	cache.Set("name", "rajab", 1)
	cache.Set("age", "17", 2)
	cache.Set("job", "swe", 2)
	time.Sleep(1 * time.Second)
	cache.Get("name")
	t.Logf("Length: %d\n", cache.keys.LengthC)
	assert.Equal(t, uint32(2), cache.keys.LengthC)
}

func TestLRUSetLengthAlreadyPresent(t *testing.T) {
	const C = 3
	cache := CreateLRUCache(C)

	cache.Set("name", "rajab", 1)
	cache.Set("age", "17", 2)
	cache.Set("name", "rahul", 2)
	t.Logf("Length: %d\n", cache.keys.LengthC)
	assert.Equal(t, uint32(2), cache.keys.LengthC)
}

func TestLRUSetLengthFull(t *testing.T) {
	const C = 3
	cache := CreateLRUCache(C)

	cache.Set("name", "rajab", -1)
	cache.Set("age", "17", -1)
	cache.Set("job", "swe", -1)
	cache.Set("passion", "programming", -1)
	t.Logf("Length: %d\n", cache.keys.LengthC)
	assert.Equal(t, uint32(3), cache.keys.LengthC)
}

func TestLRUSetLengthNotFull(t *testing.T) {
	const C = 3
	cache := CreateLRUCache(C)

	cache.Set("name", "rajab", -1)
	cache.Set("age", "17", -1)
	t.Logf("Length: %d\n", cache.keys.LengthC)
	assert.Equal(t, uint32(2), cache.keys.LengthC)
}

// Get tests
// 1. No TTL length test?
func TestLRUGetLengthNoTTL(t *testing.T) {
	const C = 3
	cache := CreateLRUCache(C)

	cache.Set("name", "rajab", -1)
	cache.Set("age", "17", -1)
	cache.Set("job", "swe", -1)
	cache.Get("name")
	cache.Get("age")
	t.Logf("Length: %d\n", cache.keys.LengthC)
	assert.Equal(t, uint32(3), cache.keys.LengthC)
}

// 2. TTL length test?
func TestLRUGetLengthTTL(t *testing.T) {
	const C = 3
	cache := CreateLRUCache(C)

	cache.Set("name", "rajab", 1*time.Second)
	cache.Set("age", "17", 2*time.Second)
	cache.Set("job", "swe", 2*time.Second)
	assert.Equal(t, uint32(3), cache.keys.LengthC)
	time.Sleep(1000 * time.Millisecond)
	cache.Get("name")
	assert.Equal(t, uint32(2), cache.keys.LengthC)
	time.Sleep(2000 * time.Millisecond)
	cache.Get("age")
	cache.Get("job")
	assert.Equal(t, uint32(0), cache.keys.LengthC)
	t.Logf("Length: %d\n", cache.keys.LengthC)
}

// 3. alread present length test?
func TestLRUGetLengthAlreadyPresent(t *testing.T) {
	const C = 3
	cache := CreateLRUCache(C)

	cache.Set("name", "rajab", 2*time.Second)
	cache.Set("age", "17", 2*time.Second)
	cache.Get("name")
	cache.Get("name")
	t.Logf("Length: %d\n", cache.keys.LengthC)
	assert.Equal(t, uint32(2), cache.keys.LengthC)
}

// 4. full length test?
func TestLRUGetLengthFull(t *testing.T) {
	const C = 3
	cache := CreateLRUCache(C)

	cache.Set("name", "rajab", -1)
	cache.Set("age", "17", -1)
	cache.Set("job", "swe", -1)
	cache.Get("name")
	t.Logf("Length: %d\n", cache.keys.LengthC)
	assert.Equal(t, uint32(3), cache.keys.LengthC)
}
