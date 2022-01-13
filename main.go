package main

/*
set https_proxy=socks5://127.0.0.1:10002
*/

import (
	"fmt"
)

func main() {
	// nums := []int { 2, 3, 1, 0, 2, 5, 3}
	// nums := []int { 0, 1, 2, 3, 4, 5, 4}
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 14, 15}
	fmt.Println(findRepeatNumber(nums))
}

// 输入：
// [2, 3, 1, 0, 2, 5, 3]
// 输出：2 或 3
func findRepeatNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	i := 0
	j := len(nums)
	h := (i + j) / 2
	for i != j {

		fmt.Println(i, h, j)
		if repeatIn(nums, i, h-1) {
			j = h - 1
			h = (i + j) / 2
		} else {
			i = h
			h = (i + j) / 2
		}
	}
	fmt.Println(i, h, j)
	if repeatIn(nums, i, j) {
		return -1
	} else {
		return j
	}
}

func repeatIn(nums []int, i, j int) bool {
	count := 0
	for _, v := range nums {
		if v >= i && v <= j {
			count++
			if count > j-i+1 {
				return false // nums left is over half
			}
		}
	}
	return true // nums right is over half
}
