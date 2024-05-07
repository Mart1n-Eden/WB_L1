package main

import (
	"fmt"
	"time"
)

func main() {
	var x bool

	go func() {
		for {
			fmt.Println("Ожидание..")

			if x {
				fmt.Println("Остановка горутины")
				return
			}
		}
	}()

	time.Sleep(time.Millisecond * 500)
	x = true
	time.Sleep(time.Millisecond * 500)
}