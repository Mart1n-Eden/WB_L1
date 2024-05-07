package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Count int
}

func (c *Counter) Incrementation(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	c.Count++
}

func main() {
	var x Counter

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			x.Incrementation(&mu)
		}()
	}

	wg.Wait()

	fmt.Println(x.Count)
}
