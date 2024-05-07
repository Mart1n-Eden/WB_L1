package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		for {
			select {
			case <-ch:
				fmt.Println("Остановка горутины")
				return
			default:
				fmt.Println("Ожидание..")
			}
		}
	}(ch)

	time.Sleep(time.Millisecond * 500)
	ch <- 1

}
