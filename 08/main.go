package main

import (
	"fmt"
)

func main() {
	var x, i, v int64
	fmt.Println("Введите число:")
	fmt.Scanln(&x)
	fmt.Printf("%064b\n", x)
	fmt.Println("Введите номер бита:")
	fmt.Scanln(&i)
	fmt.Println("Какой бит установить")
	fmt.Scanln(&v)
	switch v {
	case 0:
		x &= ^0 << i
	case 1:
		x |= 1 << i
	}
	fmt.Printf("%064b", x)
}
