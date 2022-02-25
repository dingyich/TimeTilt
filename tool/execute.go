package main

import (
	"TimeTilt/tool/single"
	"fmt"
)

func main() {
	arr := []int{-55, 90, 3, 37, 45, 2, 0, 9, 9, 78, 34, -64, -9, -1, 22, 12, 21, -31, 37, 54, 10, 13, 16, 19, 22, 25, 28, 31, 34, 37}
	m := 3

	res := single.GetMaxAllignedSubset(arr, m)
	fmt.Println(arr)
	fmt.Println(res)
}
