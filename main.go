package main

import (
	"fmt"

	"github.com/Muhammed-Rajab/go-inca/pkg/inca"
)

func main() {

	cache := inca.CreateInca(&inca.IncaConfig{
		Capacity:       3,
		EvictionPolicy: inca.TYPE_LRU,
	})

	cache.Memory.Set("name", "Rajab", -1)
	cache.Memory.Set("age", "17", -1)
	cache.Memory.Set("job", "swe", -1)
	cache.Memory.Set("lol", "gege", -1)
	cache.Memory.Get("job")
	fmt.Println(cache.Memory.Priorities())
}
