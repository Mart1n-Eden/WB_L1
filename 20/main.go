package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()

	str := strings.Fields(s)

	for i := len(str) - 1; i >= 0; i-- {
		fmt.Printf("%s ", str[i])
	}
}
