package main

import (
	"fmt"
	"github.com/shomali11/maps"
)

func main() {
	concurrentMap := maps.NewShardedConcurrentMap()
	concurrentMap.Set("name", "Raed Shomali")

	fmt.Println(concurrentMap.Contains("name"))
	fmt.Println(concurrentMap.Get("name"))
	fmt.Println(concurrentMap.Size())
	fmt.Println(concurrentMap.IsEmpty())
	fmt.Println(concurrentMap.Keys())

	concurrentMap.Remove("name")

	fmt.Println(concurrentMap.Contains("name"))
	fmt.Println(concurrentMap.Get("name"))
	fmt.Println(concurrentMap.Size())
	fmt.Println(concurrentMap.IsEmpty())
	fmt.Println(concurrentMap.Keys())

	concurrentMap.Set("name", "Raed Shomali")
	concurrentMap.Clear()

	fmt.Println(concurrentMap.Contains("name"))
	fmt.Println(concurrentMap.Get("name"))
	fmt.Println(concurrentMap.Size())
	fmt.Println(concurrentMap.IsEmpty())
	fmt.Println(concurrentMap.Keys())
}
