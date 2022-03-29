package main

import (
	"fmt"
	"sort"
)

func main() {
	for {
		n := 0
		t1, _ := fmt.Scan(&n)
		if t1 == 0 {
			return
		} else {
			num := make([]int, n)
			for i := 0; i < n; i++ {
				fmt.Scan(&num[i])
			}
			result(num)
		}
	}
}

func result(num []int) {
	res := 0
	l := len(num)
	for i := 0; i < len(num); i++ {
		// res += num[i]
		res := test(num[i:])
	}

	fmt.Println(res)
}
func test(num []int) int {
	res := 0
	if len(num) == 0 {
		return res
	}
	mid := num[0]
	res += mid
	min, max := mid
	for i := 1; i+1 < len(num); i += 2 {
		a, b := num[i], num[i+1]
		if a > b {
			a, b = b, a
		}
		if a <= min && b >= max {
			res+=mid
		} else if a < min && b < min {
			min = 
		} else if a > max && b > max {

		}
	}

	return num[len(num)/2]
}


//中位数
// 时间限制： 3000MS
// 内存限制： 589824KB
// 题目描述：
// 小团很喜欢中位数。现在给定一个序列，若其长度为奇数，那么其中位数是将序列中的数从小到大排序后位于正中间位置的数；若其长度为偶数，那么其中位数是将序列中的数从小到大排序后位于最中间的两个数的平均值。

// 现在给你一个长度为n的序列，小团想知道所有长度为奇数的区间的中位数之和为多少。

// 输入描述
// 第一行一个正整数n，表示序列中有n个数。

// 接下来一行n个空格隔开的正整数a1,a2,…an表示序列中n个数的值。

// 1<=n<=2000, 1<=ai<=100000，保证ai互不相同。

// 输出描述
// 一行一个正整数，表示给定序列中所有长度为奇数的区间的中位数之和。

// 样例输入
// 4
// 2 3 1 4
// 样例输出
// 15

// 提示
// 样例解释

// 长度为奇数的区间有[2], [3], [1], [4], [2 3 1], [3 1 4]

// 答案为2+3+1+4+2+3=15
