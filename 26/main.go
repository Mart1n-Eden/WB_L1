package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Println("Введите строку")

	fmt.Scanln(&s)
	
	m := make(map[byte]bool)

	for i := range s {
		if m[s[i]] {
			fmt.Println("false")
			return
		}
		m[s[i]] = true
	}
	
	fmt.Println("true")
}