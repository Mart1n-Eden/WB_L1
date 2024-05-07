package main

import (
	"fmt"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := make(map[int][]float64)

	for i := range arr {
		v := int(arr[i]) / 10

		m[v*10] = append(m[v*10], arr[i])
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
