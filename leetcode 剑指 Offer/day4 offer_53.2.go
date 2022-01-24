/*
题目描述：  剑指 Offer 53 - II. 0～n-1中缺失的数字
难度：  简单
编写日期：   2022 年 1 月 19 日 晚20：00
重写日期：  
*/
/*题目：  
    一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

    示例 1:

    输入: [0,1,3]
    输出: 2
    示例 2:

    输入: [0,1,2,3,4,5,6,7,9]
    输出: 8

    限制：

    1 <= 数组长度 <= 10000
*/
/*  思路
    [0, 1, 2, 3, 4, 5, 6, 7,   8，  9]   len = 10
    [0, 1, 2, 3, 4, 5, 6, 7,   9，  10]

    第一次：二分查找 nums[i-1] == i-1 && nums[i] ！= i 的

*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println("T")
}

// 默认代码模版
func missingNumber(nums []int) int {
    for i, j := 0, len(nums)-1; i <= j; {
        mid := (i + j) / 2
        if nums[mid] == mid && ( mid + 1 == len(nums) || nums[mid + 1] != mid + 1){
            return mid + 1
        } else if nums[mid] == mid {
            i = mid + 1
        } else {
            j = mid - 1
        }
    }
    return 0        //数组 num[0] 第一个数 0 不在数组中的情况
}
