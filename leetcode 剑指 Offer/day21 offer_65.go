/*
题目描述：  剑指 Offer 65. 不用加减乘除做加法
难度：  中等
编写日期：   2022 年 2 月 13 日 晚 19：00
重写日期：  
*/
/*题目：  

	写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。

    示例:

    输入: a = 1, b = 1
    输出: 2

    提示：

    a, b 均可能是负数或 0
    结果不会溢出 32 位整数

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
func add(a int, b int) int {
    flag := 0
    c := 0
    for i := 0; i < 64; i++ {
        abit := a & (1 << i)
        bbit := b & (1 << i)
        c = c | ( abit ^ bbit ^ flag)
        if abit & bbit > 0 || abit & flag > 0 || flag & bbit > 0 {
            flag = 1 << (i+1)
        } else {
            flag = 0
        }
    }
    return  c
}