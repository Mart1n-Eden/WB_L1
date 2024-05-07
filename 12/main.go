package main

import (
	"fmt"
)

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]bool)

	for i := range arr {
		m[arr[i]] = true
	}

	fmt.Println(m)
}
