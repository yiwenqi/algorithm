package main

import (
	"algorithm/src/base"
	"fmt"
)

func main() {
	//array := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	// 考虑边界情况
	array := []int{3, 0, -2, -1, 1, 2}
	a := base.ThreeSum(array)
	fmt.Println(a)
}
