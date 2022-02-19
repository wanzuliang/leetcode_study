/*
题目描述：  剑指 Offer 59 - I. 滑动窗口的最大值
难度：  困难
编写日期：   2022 年 2 月 18 日 晚 20：00
重写日期：  
*/
/*题目：  

	给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

	示例:

	输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
	输出: [3,3,5,5,6,7] 
	解释: 

	  滑动窗口的位置                最大值
	---------------               -----
	[1  3  -1] -3  5  3  6  7       3
	 1 [3  -1  -3] 5  3  6  7       3
	 1  3 [-1  -3  5] 3  6  7       5
	 1  3  -1 [-3  5  3] 6  7       5
	 1  3  -1  -3 [5  3  6] 7       6
	 1  3  -1  -3  5 [3  6  7]      7
	 

	提示：

	你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

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
func maxSlidingWindow(nums []int, k int) []int {
    if nums == nil {
        return nil
    }
    if len(nums) == 0 || k == 1 {
        return nums
    }
    res := make([]int, len(nums) - k + 1)
    max := 0
    max_index := -1
    for i := 0; i + k <= len(nums); i++ {
        if i == 0 {
            max_index := findMax(nums[i:i + k])
            max = nums[i + max_index]
            res[i] = max
        }
        
        if i <= max_index && nums[i+k-1] < max {
            res[i] = max
            continue
        } else if nums[i+k-1] >= max {
            max_index = i+k-1
            max = nums[i+k-1]
            res[i] = max
        } else {
            max_index := findMax(nums[i:i + k])
            max = nums[i + max_index]
            res[i] = max
        }
    }
    return res
}

func findMax(nums []int) int {
    max := nums[0]
    res := 0
    for index, v := range nums {
        if v >= max {
            res = index
            max = v
        }
    }
    return res
}