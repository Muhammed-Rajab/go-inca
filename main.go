package main

import (
	"github.com/Muhammed-Rajab/go-inca/pkg/inca"
)

func main() {

	cache := inca.CreateInca(&inca.IncaConfig{
		Capacity:       3,
		EvictionPolicy: inca.TYPE_LRU,
	})

	cache.Run(`GET name`)
	cache.Run(`SET name "Rajab"`)
	cache.Run(`GET name`)
	cache.Run(`EXPIRE name 10`)
	cache.Run(`EXPIRE name -1`)
	cache.Run(`KEYS`)
	cache.Run(`TTL name`)
}
