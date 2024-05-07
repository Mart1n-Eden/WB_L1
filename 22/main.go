package main

import (
	"fmt"
	"math/big"
)

// 2^20 = 1048576
// 2^21 = 2097152
// 2^22 = 4194304

func main() {
	var x, y int64
	var op string
	
	fmt.Println("Введите 2 числа")
	fmt.Scan(&x, &y)
	fmt.Println("Введите операцию")
	fmt.Scan(&op)

	bx := big.NewInt(x)
	by := big.NewInt(y)

	Calc(bx,by,op)
}

func Calc(x,y *big.Int, op string) {
	switch op {
	case "+":
		fmt.Println(x.Add(x,y))
	case "*":
		fmt.Println(x.Mul(x,y))
	case "-":
		fmt.Println(x.Sub(x,y))
	case "/":
		if y.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("Деление на 0")
		} else {
			fmt.Println(x.Div(x,y))
		}
	}
}