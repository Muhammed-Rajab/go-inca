package main

import (
	"fmt"

	"github.com/Muhammed-Rajab/go-inca/pkg/inca"
)

func main() {

	cache := inca.CreateInca(&inca.IncaConfig{
		Capacity:       10,
		EvictionPolicy: inca.TYPE_LRU,
	})

	cache.Memory.Set("name", "Rajab")
	fmt.Println(cache.Memory.Priorities())
	fmt.Println(cache.Memory.Get("name"))

	cache.Memory.Set("name", "Jamie")
	fmt.Println(cache.Memory.Priorities())
	fmt.Println(cache.Memory.Get("name"))
}
