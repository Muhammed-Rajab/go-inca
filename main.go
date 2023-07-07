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
	cache.Memory.Set("address", "house", -1)
	cache.Memory.Set("phno", "8098123122", -1)
	cache.Memory.Set("age", "boob inspector", -1)
	fmt.Println(cache.Memory.Priorities())
}
