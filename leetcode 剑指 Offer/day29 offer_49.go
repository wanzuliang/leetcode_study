/*
题目描述：  剑指 Offer 49. 丑数
难度：  中等
编写日期：   2022 年 2 月 24 日 晚 17：00
重写日期：  	2022 年 2 月 27 日 晚 17：00
*/
/*题目：  

	我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

	示例:

	输入: n = 10
	输出: 12
	解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
	
	说明:  

	1 是丑数。
	n 不超过1690。

*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
	res := nthUglyNumber(1690)
    fmt.Println( res )
}

// 默认代码模版
func nthUglyNumber(n int) int {
    list := []int{1}
    a2 := 0
    a3 := 0
    a5 := 0
    for rank := 1; rank < n; {
        if list[a2]*2 <= list[a3]*3 && list[a2]*2 <= list[a5]*5 {
            if list[a2]*2 > list[len(list)-1] {
                list = append(list, list[a2]*2)
                rank++
            }
            a2++
            continue
        }
        if list[a3]*3 <= list[a2]*2 && list[a3]*3 <= list[a5]*5 {
            if list[a3]*3 > list[len(list)-1] {
                list = append(list, list[a3]*3)
                rank++
            }
            a3++
            continue
        }
        if list[a5]*5 <= list[a3]*3 && list[a5]*5 <= list[a2]*2 {
            if list[a5]*5 > list[len(list)-1] {
                list = append(list, list[a5]*5)
                rank++
            }
			a5++
            continue
        }
    }
    return list[len(list)-1]
}