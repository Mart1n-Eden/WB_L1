package main

import (
	"fmt"
	"strings"
)
// Проблема в том, что выделяется память под 1000 элементов, а используется
// только 100, при том что все 1000 хранятся в куче все время работы программы
var justString string

func createHugeString(n int) string {
	return strings.Repeat("@", n)
}

func someFunc() {
	v := createHugeString(2 << 10)
	tmp := make([]byte, 100)
	copy(tmp, v[:100])
	justString = string(tmp)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
