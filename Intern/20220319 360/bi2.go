package main

import (
	"fmt"
	// "math"
)

/*
	题目大概为：
	某人射击，2 发机会。可能没中，可能中 1 发， 也可能中 2 发。
	有 n 个目标物，击中第 i 个目标 得 Ai 分，如果击中两发 i 和 j ，则取 (Ai | Aj) 为结果分数
	
	输入整数 n
	再输入长度为 n 的数组，代表每个目标的分数

	输出 所有可能结果加在一起的总分数

	例如：
	输入：
	>> 3
	>> 1 2 3
	输出：
	>> 15
*/

func main() {
	n := 0
	t1, _ := fmt.Scan(&n)
	if t1 == 0 {
		return
	} else {
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		fmt.Printf("%d", result(a))
	}
}

func result(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]						// 只有 1 发命中
		for j := i+1; j < len(nums); j++ {
			res += nums[i] | nums[j]		// 		2 发都命中
		}
	}
	return res
}