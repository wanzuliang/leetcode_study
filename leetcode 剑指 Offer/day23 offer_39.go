/*
题目描述：  剑指 Offer 39. 数组中出现次数超过一半的数字
难度：  简单
编写日期：   2022 年 2 月 15 日 晚 16：00
重写日期：  
*/
/*题目：  

	数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

    你可以假设数组是非空的，并且给定的数组总是存在多数元素。


    示例 1:

    输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
    输出: 2

    限制：
    1 <= 数组长度 <= 50000


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
// func majorityElement(nums []int) int {
//     for i := 0; i <= len(nums)/2; i++ {
//         min := i
//         for j := i + 1; j < len(nums); j++ {
//             if nums[j] < nums[min] {
//                 min = j
//             }
//         }
//         if min != i {
//             temp := nums[i]
//             nums[i] = nums[min]
//             nums[min] = temp
//         }
//     }
//     return nums[len(nums)/2]
// }
func majorityElement(nums []int) int {
    res, bit := 0, 0
    for i := 0; i < len(nums); i++ {
        if bit == 0 {
            res = nums[i]
            bit++
        } else {
            if res == nums[i] {
                bit++
            } else {
                bit--
            }
        }
    }
    return res
}