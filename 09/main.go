package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	input := make(chan int)
	output := make(chan int)

	go func() {
		defer close(output)
		for v := range input {
			output <- v * 2
		}
	}()

	go func() {
		defer close(input)
		for i := range arr {
			input <- arr[i]
		}
	}()

	for v := range output {
		fmt.Println(v)
	}
}
