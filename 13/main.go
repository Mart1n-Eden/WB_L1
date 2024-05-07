package main

import (
	"fmt"
)

func main() {
	x,y := 5,6

	fmt.Println(x,y)

	x,y = y,x

	fmt.Println(x,y)
}