package main

import (
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
		node := dll.CreateNode(key, val, nil, nil)
		cache.data[key] = node
		cache.keys.Prepend(node)
		return
	}

	// Case 2 -> Cache is full
	node := dll.CreateNode(key, val, nil, nil)
	popped := cache.keys.Pop()
	delete(cache.data, popped.Key)
	cache.data[key] = node
	cache.keys.Prepend(node)
}

// Method to Get value from cache
func (cache *LRUCache) Get(key string) (string, bool) {

	// O/P: (val, isPresent)

	if node := cache.data[key]; node == nil {
		// Case 1 -> Key is not present
		return "", false
	} else {
		// Case 2 -> Key is present

		// Removed node from the equation
		// if node is head
		if node == cache.keys.Head() {
			return node.Val, true
		}

		// if node is tail
		if node.Next == nil {
			node.Prev.Next = nil
			cache.keys.Prepend(node)
			return node.Val, true
		}

		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		// prepend node
		cache.keys.Prepend(node)
		return node.Val, true
	}
}

// Method to check if the capacity is reached
// Optimized way: keep track of capacity using another variable in the struct
func (cache *LRUCache) IsFull() bool {
	return cache.keys.Length() == cache.capacity
}

func main() {
	cache := CreateLRUCache(5)
	cache.Set("a", "1")
	cache.Set("b", "2")
	cache.Set("c", "3")
	cache.Set("d", "4")
	cache.Set("e", "5")

	cache.Set("f", "6")
	cache.keys.Display()
}
