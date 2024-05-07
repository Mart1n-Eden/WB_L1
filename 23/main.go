package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	var i int
	fmt.Scanln(&i)

	arr = append(arr[:i], arr[i+1:]...)

	fmt.Println(arr)
}
