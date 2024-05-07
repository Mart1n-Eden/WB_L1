package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Scanln(&s)

	str := []rune(s)

	for i := len(str)-1; i >= 0; i-- {
		fmt.Printf("%c",str[i])
	}
}
