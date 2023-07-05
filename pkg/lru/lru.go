package lru

import (
	dll "github.com/Muhammed-Rajab/go-inca/pkg/dll"
)

type LRUCache struct {
	capacity uint32
	data     map[string]string
	keys     *dll.DoublyLinkedList
}
