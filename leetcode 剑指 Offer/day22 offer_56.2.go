/*
题目描述：  剑指 Offer 56 - II. 数组中数字出现的次数 II
难度：  中等
编写日期：   2022 年 2 月 14 日 晚 20：00
重写日期：  
*/
/*题目：  

	在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

    示例 1：

    输入：nums = [3,4,3,3]
    输出：4
    示例 2：

    输入：nums = [9,1,7,9,7,9,7]
    输出：1

    限制：

    1 <= nums.length <= 10000
    1 <= nums[i] < 2^31

*/
/*  
        数电    卡诺图
*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( "T")
}

// 默认代码模版
func singleNumber(nums []int) int {
    a, b := 0, 0
    for _, c := range nums {
        ta := a
        a = (b & c) | (a & ^c)
        b = (b & ^c) | (^ta & ^b & c)
    }
    return b
}