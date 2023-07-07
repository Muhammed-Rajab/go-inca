package main

import (
	"fmt"
	"time"

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
	cache.Memory.ExpireTTL("name", 1*time.Millisecond)

	time.Sleep(1 * time.Second)
	fmt.Println(cache.Memory.GetTTL("age"))

	fmt.Println(cache.Memory.Get("name"))
	fmt.Println(cache.Memory.Priorities())
}
