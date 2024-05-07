package main

import (
	"math/rand"
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for k := 0; k < 10; k++{
		wg.Add(1)
		
		v := rand.Int()
		go func(k int, v int) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			m[k] = v
		}(k,v)
	}

	wg.Wait()

	for k,v := range m {
		fmt.Println(k, ":", v)
	}
}