package inca

import (
	"time"

	"github.com/Muhammed-Rajab/go-inca/pkg/lru"
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

	// Helpers
	IsFull() bool
	Priorities() []string

	// TTL
	GetTTL(key string) (time.Duration, bool)
	ExpireTTL(key string, duration time.Duration) bool
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
