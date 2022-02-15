/*
题目描述：  剑指 Offer 66. 构建乘积数组
难度：  中等
编写日期：   2022 年 2 月 15 日 晚 16：00
重写日期：  
*/
/*题目：  

	给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

    示例:

    输入: [1,2,3,4,5]
    输出: [120,60,40,30,24]
     

    提示：

    所有元素乘积之和不会溢出 32 位整数
    a.length <= 100000
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
func constructArr(a []int) []int {
    res := make([]int, len(a))
    for i, t := 0, 1; i < len(a); i++ {
        res[i] = t
        t *= a[i]
    }
    for i, t := len(a)-1, 1; i >= 0 ; i-- {
        res[i] *= t
        t *= a[i]
    }
    return res
}