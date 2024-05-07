package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Остановка горутины")
				ch <- 1
				return
			default:
				fmt.Println("Ожидание..")
			}
		}
	}(ctx)

	time.Sleep(time.Millisecond * 500)
	cancel()

	<-ch
}
