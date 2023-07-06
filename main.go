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

	cache.Memory.Set("name", "Rajab")
	cache.Memory.Set("age", "17")
	cache.Memory.Set("job", "swe")

	fmt.Println(cache.Memory.Get("name"))
	fmt.Println(cache.Memory.Priorities())
}
