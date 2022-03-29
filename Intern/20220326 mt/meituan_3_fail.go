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
	m := make([]int, 7)
	t := make([][]int, 0)
	for i := 0; i < len(num); i++ {
		if num[i] < 0 {
			switch num[i] % 7 {
			case 1:
				if num[i] > m[1] {
					m[1] = num[i]
				}
			case 2:
				if num[i] > m[2] {
					m[2] = num[i]
				}
			case 3:
				if num[i] > m[3] {
					m[3] = num[i]
				}
			case 4:
				if num[i] > m[4] {
					m[4] = num[i]
				}
			case 5:
				if num[i] > m[5] {
					m[5] = num[i]
				}
			case 6:
				if num[i] > m[6] {
					m[6] = num[i]
				}
			}
		} else {
			switch num[i] % 7 {
			case 0:
				res+=num[i]
			case 1:
				t[1] = append(t[1], num[i])
			case 2:
				t[2] = append(t[2], num[i])
			case 3:
				t[3] = append(t[3], num[i])
			case 4:
				t[4] = append(t[4], num[i])
			case 5:
				t[5] = append(t[5], num[i])
			case 6:
				t[6] = append(t[6], num[i])
			}
		}
	}
	mt := make([]int, 7)
	for i := 0; i < len(t[6]); i++{
		if 
	}
	fmt.Println(res)
}

// 时间限制： 3000MS
// 内存限制： 589824KB
// 题目描述：
// 小美喜欢7的倍数。桌面上有一些卡片，每张卡片上都印有一个数字，小美想从中挑选一些卡片，使得卡片上的数字之和最大，由于小美很喜欢7的倍数，她同时还希望挑选出的卡片的数字之和是7的倍数，请问她能挑选出的最大数字之和是多少？（注意，小美也可以一张卡片都不挑选）

// 输入描述
// 第一行是一个正整数n，表示桌上有n张写有数字的卡片。

// 第二行有n个空格隔开的整数a1,a2,…,an，其中ai表示桌上第i张卡片上所写的数字。

// 1<=n<=50000, |ai|<=3000

// 输出描述
// 一行一个整数，表示小美能挑选出的最大数字之和（且和为7的倍数）。
