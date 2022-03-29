package main

import (
	"fmt"
)

func main() {
	for {
		n := 0
		t1, _ := fmt.Scan(&n)
		if t1 == 0 {
			return
		} else {
			num := make([]int64, n)
			for i := 0; i < n; i++ {
				fmt.Scan(&num[i])
			}
			result(num)
		}
	}
}

func result(num []int64) {
	if len(num) == 2 {
		fmt.Println(num[1] - num[0])
		return
	}
	a, b := num[0], num[len(num)-1]
	res := b - a
	for i := 1; i < len(num)-1; i++ {
		temp := a + b - 2*num[i]
		if temp < 0 {
			temp = int64(-1) * temp
		}
		if temp < res {
			res = temp
		}
	}
	fmt.Println(res)
}

// 时间限制： 3000MS
// 内存限制： 589824KB
// 题目描述：
// 数轴上有n个点，从左到右编号分别为1,2,…,n。

// 小美在1号点，小团在n号点，现在要选择一个点作为他们会合的地点，他们期望选择的点能让小美和小团到达会合点的距离差值尽量小。

// 你的任务是输出最小的距离差。

// 输入描述
// 第1行是一个正整数n，表示数轴上有n个点。

// 第2行是n个空格隔开的正整数a1,a2,…,an，第 i 个数表示第 i 个点的坐标。

// 2<=n<=100000，1<=ai<=200，保证a1<=a2<=…<=an

// 输出描述
// 输出一个整数，表示最小距离差值。
