package main

import (
	"algorithm/src/sort_m"
	"fmt"
)

func main() {
	//array := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	// 考虑边界情况
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := sort_m.QuickSort(array)
	fmt.Println(a)
}
