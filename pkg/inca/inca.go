package inca

import (
	"fmt"
	"strings"
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

func (cache *Inca) Run(query string) (string, error) {
	parsed, err := queryparser.Parse(query)

	// Handle parser error
	if err != nil {
		return "", err
	}

	// Run command based on it
	if parsed.CommandType == "GET" {
		// Return val if present, else return error
		if val, present := cache.Memory.Get(parsed.Args.Key); present {
			return val, nil
		} else {
			return "", fmt.Errorf("GET error: key '%s' not present", parsed.Args.Key)
		}
	} else if parsed.CommandType == "DELETE" {
		// If deleted, then return DONE, else return error
		if done := cache.Memory.Delete(parsed.Args.Key); !done {
			return "", fmt.Errorf("GET error: key '%s' not present", parsed.Args.Key)
		} else {
			return "DONE", nil
		}
	} else if parsed.CommandType == "CLEAR" {
		cache.Memory.Clear()
		return "DONE", nil
	} else if parsed.CommandType == "TTL" {
		if duration, present := cache.Memory.GetTTL(parsed.Args.Key); !present {
			return "", fmt.Errorf("GET error: key '%s' not present", parsed.Args.Key)
		} else {
			return duration.String(), nil
		}
	} else if parsed.CommandType == "SET" {
		// if ttl is "-1", then set it to -1 ns
		var duration time.Duration = -1
		if parsed.Args.TTL != "-1" {
			duration, _ = time.ParseDuration(parsed.Args.TTL + "s")
		}
		cache.Memory.Set(parsed.Args.Key, parsed.Args.Val, duration)
		return "DONE", nil
	} else if parsed.CommandType == "EXPIRE" {
		var duration time.Duration = -1
		if parsed.Args.TTL != "-1" {
			duration, _ = time.ParseDuration(parsed.Args.TTL + "s")
		}
		cache.Memory.ExpireTTL(parsed.Args.Key, duration)
		return "DONE", nil
	} else if parsed.CommandType == "KEYS" {
		keys := cache.Memory.Priorities()
		quotes := "\""
		if len(keys) == 0 {
			quotes = ""
		}
		return fmt.Sprintf("[%s%s%s]", quotes, strings.Join(keys, "\", \""), quotes), nil
	}

	return "", fmt.Errorf("syntax error: invalid query `%s`", query)
}
