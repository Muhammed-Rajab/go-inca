package main

import (
	"fmt"

	"github.com/Muhammed-Rajab/go-inca/pkg/inca"
)

func main() {

	// Capacity of the cache
	const CACHE_CAPACITY = 10

	// Always call this function to create a new cache
	cache := inca.CreateInca(&inca.IncaConfig{
		Capacity:       CACHE_CAPACITY,
		EvictionPolicy: inca.TYPE_LRU,
	})

	// Your code goes here.....
	cache.Run(`SET name "Markus"`)
	if name, err := cache.Run(`GET namej`); err != nil {
		panic(err)
	} else {
		fmt.Printf("The name is %s\n", name)
	}
}
