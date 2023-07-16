package inca

import (
	"fmt"
	"time"

	"github.com/Muhammed-Rajab/go-inca/pkg/dll"
	"github.com/Muhammed-Rajab/go-inca/pkg/lru"
	queryparser "github.com/Muhammed-Rajab/go-inca/pkg/query_parser"
)

// Configs
type CacheType int

const (
	TYPE_LRU CacheType = iota
	TYPE_LFU
)

type IncaConfig struct {
	Capacity       uint32
	EvictionPolicy CacheType
}

// Inca
type Cache interface {
	// CRUD
	Get(key string) (string, bool)
	Set(key, val string, ttl time.Duration)
	Delete(key string) bool
	Clear()

	// Helpers
	IsFull() bool
	Priorities() []string

	// TTL
	GetTTL(key string) (time.Duration, bool)
	ExpireTTL(key string, duration time.Duration) bool

	// Testing
	Keys() *dll.DoublyLinkedList
	Data() map[string]*dll.Node
}

type Inca struct {
	Config IncaConfig
	Memory Cache
}

func CreateInca(config *IncaConfig) *Inca {
	inca := &Inca{}

	// Change this code to support both LRU and LFU
	inca.Config = *config
	inca.Memory = lru.CreateLRUCache(config.Capacity)
	return inca
}

func (cache *Inca) Run(query string) {
	parsed, err := queryparser.Parse(query)

	// Handle parser error
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run command based on it
	if parsed.CommandType == "GET" {
		if val, present := cache.Memory.Get(parsed.Args.Key); present {
			fmt.Printf("%s\n", val)
		} else {
			fmt.Printf("ERROR: Not present\n")
		}
	} else if parsed.CommandType == "DELETE" {
		if done := cache.Memory.Delete(parsed.Args.Key); !done {
			fmt.Printf("ERROR: Not present\n")
		} else {
			fmt.Printf("DONE\n")
		}
	} else if parsed.CommandType == "CLEAR" {
		cache.Memory.Clear()
		fmt.Printf("DONE\n")
	} else if parsed.CommandType == "TTL" {
		if duration, present := cache.Memory.GetTTL(parsed.Args.Key); !present {
			fmt.Printf("ERROR: Not present\n")
		} else {
			fmt.Printf("%s\n", duration)
		}
	}
}
