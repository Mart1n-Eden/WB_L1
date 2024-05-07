package main

import (
	"fmt"
	"sort"
)

func binSearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{3, 5, 7, 1, 45, 7, 3, 14, 87, 9, 44}

	sort.Ints(arr)
	fmt.Println(arr)

	fmt.Println(binSearch(arr, 14))
}
