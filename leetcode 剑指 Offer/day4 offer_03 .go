/*
题目描述：  剑指 Offer 03. 数组中重复的数字
难度：  简单
编写日期：   2022 年 1 月 18 日 晚20：30
重写日期：  
*/
/*题目：  
    找出数组中重复的数字。
    在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
    
    示例 1：
    输入：
    [2, 3, 1, 0, 2, 5, 3]
    输出：2 或 3 
    
    限制：
    2 <= n <= 100000
*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println("T")
}

// 默认代码模版
// func findRepeatNumber(nums []int) int {
//     for i := 0; i < len(nums); i++ {
//         if nums[i] == i {
//             continue
//         } else if nums[i] != nums[nums[i]]{
//             temp := nums[nums[i]]
//             nums[nums[i]] = nums[i]
//             nums[i] = temp
//         } else {
//             return nums[i]
//         }
//     }
//     return -1
// }
// [3, 4, 2, 0, 0, 1]
// 输出：
// -1
// 预期结果：
// 0


func findRepeatNumber(nums []int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == i {
            continue
        } else if nums[i] != nums[nums[i]]{
            temp := nums[nums[i]]
            nums[nums[i]] = nums[i]
            nums[i] = temp
            if nums[i] != i {   // 置换后， 置换至 i 处的值需要再判断一次
                i--             // 故 i-- 回退一次与 i++ 抵消
            }
        } else {
            return nums[i]
        }
    }
    return -1
}