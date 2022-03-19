package main

import (
	"fmt"
	// "math"
)

/*
	题目大概为：
	求出字符串 str 中的子串重复最多的次数，输出整数。

	例如：
	输入：
	>> aba
	输出：
	>> 2
	因为 a 重复了两次
*/

func main() {
	var str string
	t1, _ := fmt.Scan(&str)
	if t1 == 0 {
		return
	} else {
		fmt.Printf("%d", result(str))
	}
}

func result(str string) int {		// 看作是找单个字符出现最多的次数，直接遍历
	if len(str) == 0 {
		return 0
	}
	list := [26]int{}
	for i := 0; i < len(str); i++ {	// 计数
		list[str[i]-'a']++
	}
	res := 0
	for i := 0; i < 26; i++ {		// 找最大值
		if list[i] > res {
			res = list[i]
		}
	}
	return res
}