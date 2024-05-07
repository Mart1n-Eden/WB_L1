package main

import (
	"fmt"
)

func main() {
	arr1 := []int{4,3,6,8,2,1,13,4}
	arr2 := []int{9,0,112,3,2,7,8}
	res := make([]int,0)

	m := make(map[int]bool)

	for i := range arr1 {
		m[arr1[i]] = true
	}
	
	for i := range arr2 {
		if m[arr2[i]] {
			res = append(res, arr2[i])
		}
	}

	fmt.Println(res)

}