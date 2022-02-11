/*
题目描述：  剑指 Offer 64. 求1+2+…+n
难度：  中等
编写日期：   2022 年 2 月 11 日 晚 19：00
重写日期：  
*/
/*题目：  

	求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

	示例 1：

	输入: n = 3
	输出: 6
	示例 2：

	输入: n = 9
	输出: 45
	 

	限制：

	1 <= n <= 10000

*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( "T")
}

// 默认代码模版
func sumNums(n int) int {
    return (int( math.Pow(float64(n), 2) ) + n ) >> 1
}