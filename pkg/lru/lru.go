package lru

import (
	"time"

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
func (cache *LRUCache) Set(key, val string, ttl time.Duration) {

	// Potential bug, not updating tail node

	// Case 0 -> Key is already present and is trying to reset the value
	if cache.data[key] != nil {
		// Retrieve node
		node := cache.data[key]

		// Reset values to new
		node.Reset(key, val, ttl)

		// Move the node to the beginning of keys
		if node == cache.keys.Head() {
			// if node is head
			return
		} else if node.Next == nil {
			// if node is tail
			node.Prev.Next = nil
			cache.keys.Prepend(node)
		} else {
			node.Remove()
			cache.keys.Prepend(node)
		}
	}

	// Case 1 -> Cache is not full
	if !cache.IsFull() {
		node := dll.CreateNode(key, val, ttl, nil, nil)
		cache.data[key] = node
		cache.keys.Prepend(node)
		return
	}

	// Case 2, better implementation
	popped := cache.keys.Pop()
	popped.Reset(key, val, ttl)
	delete(cache.data, popped.Key)
	cache.data[key] = popped
	cache.keys.Prepend(popped)
}

// Method to Get value from cache
func (cache *LRUCache) Get(key string) (string, bool) {

	// O/P: (val, isPresent)
	node := cache.data[key]

	if node == nil {
		// Case 1 -> Key is not present

		return "", false
	}

	// Checking if the node has expired
	if node.TTL != -1 {
		if time.Since(node.StoredAt) >= node.TTL {
			// Remove the node
			// if head, update both node
			if node == cache.keys.Head() {

				cache.keys.HeadNode = cache.keys.HeadNode.Next
				if cache.keys.HeadNode != nil {
					cache.keys.HeadNode.Prev = nil
				}
			} else if node == cache.keys.Tail() {
				// if tail, update both node
				cache.keys.TailNode = cache.keys.TailNode.Prev
				if cache.keys.TailNode != nil {
					cache.keys.TailNode.Next = nil
				}
			} else {
				// else
				node.Prev.Next = node.Next
				node.Next.Prev = node.Prev
			}
			// Return nothing
			return "", false
		}
	}

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

// Method to Delete value from cache
func (cache *LRUCache) Delete(key string) bool {
	if node := cache.data[key]; node == nil {
		// Case 1 -> key is not present
		return cache.data[key] == nil
	} else {
		// Case 2 -> key is present

		// if head
		if node == cache.keys.Head() {
			// Remove from keys
			cache.keys.HeadNode = cache.keys.HeadNode.Next
			// Remove from data
			delete(cache.data, key)
			return true
		}
		// if tail
		if node.Next == nil {
			// Remove from keys
			node.Prev.Next = nil
			// Remove from data
			delete(cache.data, key)
			return true
		}
		// else
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		delete(cache.data, key)
		return true
	}
}

// Method to check if the capacity is reached
// Optimized way: keep track of capacity using another variable in the struct
func (cache *LRUCache) IsFull() bool {
	return cache.keys.Length() == cache.capacity
}

// Method to return the list format of key priorities
func (cache *LRUCache) Priorities() []string {
	temp := cache.keys.HeadNode
	pkeys := []string{}

	for temp != nil {
		pkeys = append(pkeys, temp.Key)
		temp = temp.Next
	}

	return pkeys
}

func (cache *LRUCache) ExpireTTL(key string, duration time.Duration) bool {
	if cache.data[key] != nil {
		cache.data[key].TTL = duration
		cache.data[key].StoredAt = time.Now()
		return true
	}
	return false
}

func (cache *LRUCache) GetTTL(key string) (time.Duration, bool) {
	node := cache.data[key]
	if node != nil {
		if node.TTL == -1 {
			return -1, true
		}
		return time.Until(node.StoredAt.Add(node.TTL)), true
	}
	return -1, false
}
