/*
题目描述：  剑指 Offer 53 - I. 在排序数组中查找数字 I
难度：  简单
编写日期：   2022 年 1 月 18 日 晚20：30
重写日期：   2022 年 1 月 19 日 晚20：00
*/
/*题目：  
    统计一个数字在排序数组中出现的次数。
    示例 1:
    输入: nums = [5,7,7,8,8,10], target = 8
    输出: 2

    示例 2:
    输入: nums = [5,7,7,8,8,10], target = 6
    输出: 0

    提示：
    0 <= nums.length <= 10^5
    -10^9 <= nums[i] <= 10^9
    nums是一个非递减数组
    -10^9 <= target <= 10^9
    
*/
/*  思路
    第一次：两端设置 i，j。向中间夹逼。 不太好写代码。。。。

    第二次：二分查找到 target 值，再向两边扩展

    第三次：两次二分法分别查找左右边界
*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println("T")
}

// 默认代码模版
func search(nums []int, target int) int {
    start, end := 0, len(nums)-1
    right, left := 0, -1
    for i, j := start, end; i <= j; {               //二分找 target 左边界
        mid := (i + j) / 2
        // 边界 i-1 < T  && i == T
        if nums[mid] == target && (mid == 0 || nums[mid-1] != target){
            right = mid
            break
        } else if nums[mid] < target {
            i = mid + 1
        } else {
            j = mid - 1
        }
    }
    for i, j := start, end; i <= j; {               //二分找 target 右边界
        mid := (i + j) / 2
        // 边界 i+1 > T  && i == T
        if nums[mid] == target && (mid == end || nums[mid+1] != target){
            left = mid
            break
        } else if nums[mid] > target {
            j = mid - 1
        } else {
            i = mid + 1
        }
    }
    return left - right + 1
}

