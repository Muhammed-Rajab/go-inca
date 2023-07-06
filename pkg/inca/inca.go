package inca

import "github.com/Muhammed-Rajab/go-inca/pkg/lru"

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
	Get(key string) (string, bool)
	Set(key, val string)
	Delete(key string) bool
	IsFull() bool
	Priorities() []string
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
