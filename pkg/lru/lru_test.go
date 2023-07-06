package lru

import (
	"strconv"
	"testing"
)

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
		cache.Set(kv, kv)
	}

	if cache.keys.Tail().Val != strconv.FormatInt(int64(2), 10) {
		t.Errorf("LRUCacheSet test failed! output: %s, expected: %s.\n", cache.keys.Tail().Val, "2")
	}

	if cache.keys.Head().Val != strconv.FormatInt(int64(C+1), 10) {
		t.Errorf("LRUCacheSet test failed! output: %s, expected: %s.\n", cache.keys.Head().Val, strconv.FormatInt(int64(C+1), 10))
	}
}
