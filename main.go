package main

import (
	"github.com/Muhammed-Rajab/go-inca/pkg/inca"
)

func main() {

	cache := inca.CreateInca(&inca.IncaConfig{
		Capacity:       10,
		EvictionPolicy: inca.TYPE_LRU,
	})
	cache.Memory.Get("hi")
}
