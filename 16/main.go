package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	base := arr[len(arr)/2]
	var less, more []int

	for i := range arr {
		if arr[i] < base {
			less = append(less, arr[i])
		} else if arr[i] > base {
			more = append(more, arr[i])
		}
	}

	less = quickSort(less)
	more = quickSort(more)

	less = append(less, base)

	return append(less, more...)
}

func main() {
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Println(arr)

	arr = quickSort(arr)
	fmt.Println(arr)
}
