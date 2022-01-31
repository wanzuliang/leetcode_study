/*
题目描述：  剑指 Offer 57. 和为s的两个数字
难度：  简单
编写日期：   2022 年 1 月 31 日 晚 20：30
重写日期：  
*/
/*题目：  
    输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

    示例 1：

    输入：nums = [2,7,11,15], target = 9
    输出：[2,7] 或者 [7,2]
    示例 2：

    输入：nums = [10,26,30,31,47,60], target = 40
    输出：[10,30] 或者 [30,10]

    限制：

    1 <= nums.length <= 10^5
    1 <= nums[i] <= 10^6

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
func twoSum(nums []int, target int) []int {
    for i, j := 0, len(nums) - 1; i < j; {
        if nums[i] + nums[j] == target {
            return []int{nums[i], nums[j]}
        }
        if  nums[i] + nums[j] < target {
            i++
        } else {
            j--
        }
    }
    return nil
}