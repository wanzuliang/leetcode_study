/*
题目描述：  剑指 Offer 56 - I. 数组中数字出现的次数
难度：  中等
编写日期：   2022 年 2 月 14 日 晚 20：00
重写日期：  
*/
/*题目：  

	一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。


    示例 1：

    输入：nums = [4,1,4,6]
    输出：[1,6] 或 [6,1]
    示例 2：

    输入：nums = [1,2,10,4,1,4,3,3]   2 * n + a + b = 28
    输出：[2,10] 或 [10,2]

    限制：

    2 <= nums.length <= 10000

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
func singleNumbers(nums []int) []int {
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum ^= nums[i]
    }
    bit := 0
    for i := 0; i < 64; i++ {
        if 1 << i & sum == 1 << i {
            bit = 1 << i
            break
        }
    }
    a := 0
    b := 0
    for j := 0; j < len(nums); j++ {
        if bit & nums[j] == bit {
            a ^= nums[j]
        } else {
            b ^= nums[j]
        }
    }
    return []int{a, b}
}