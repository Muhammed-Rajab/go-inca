package main

import (
	"fmt"

	dll "github.com/Muhammed-Rajab/go-inca/pkg/dll"
)

// ////////////////////////////
// LRUCache
// ////////////////////////////
type LRUCache struct {
	capacity uint32
	data     map[string]*dll.Node
	keys     *dll.DoublyLinkedList
}

// Function to create LRUCache of fixed capacity
func CreateLRUCache(capacity uint32) *LRUCache {
	cache := &LRUCache{}
	cache.capacity = capacity
	cache.data = map[string]*dll.Node{}
	cache.keys = dll.CreateDoublyLinkedListEmpty()
	return cache
}

// Method to Set values to cache
func (cache *LRUCache) Set(key, val string) {

	// Case 1 -> Cache is not full
	if !cache.IsFull() {
		node := dll.CreateNode(val, nil, nil)
		cache.data[key] = node
		cache.keys.Prepend(node)
		return
	}

	// Case 2 -> Cache is full
	node := dll.CreateNode(val, nil, nil)
	cache.data[key] = node
	cache.keys.Pop()
	cache.keys.Prepend(node)
}

// Method to Get value from cache
func (cache *LRUCache) Get(key string) string {
	return cache.data[key].Val
}

// Method to check if the capacity is reached
// Optimized way: keep track of capacity using another variable in the struct
func (cache *LRUCache) IsFull() bool {
	return cache.keys.Length() == cache.capacity
}

func main() {
	cache := CreateLRUCache(1)
	cache.Set("name", "Rajab")
	cache.Set("age", "17")

	fmt.Println(cache.Get("age"))
	fmt.Println(cache.Get("name"))
}
