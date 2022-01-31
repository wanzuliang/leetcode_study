/*
题目描述：  剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
难度：  简单
编写日期：   2022 年 1 月 31 日 晚 20：30
重写日期：  
*/
/*题目：  
    输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。

	示例：

	输入：nums = [1,2,3,4]
	输出：[1,3,2,4] 
	注：[3,1,2,4] 也是正确的答案之一。

	提示：

	0 <= nums.length <= 50000
	0 <= nums[i] <= 10000

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
func exchange(nums []int) []int {	// 这样直接改变数组的值也能通过leetcode，按说值传递应该改不了数组数据
    for i, j := 0, len(nums) - 1; i < j;{
        for ; nums[i] % 2 == 1 && i < j; {
            i++
            
        }
        for ; nums[j] % 2 == 0 && i < j; {
            j--
        }
        temp := nums[i]
        nums[i] = nums[j]
        nums[j] = temp
        i++
        j--
    }
    return nums
}

func exchange(nums []int) []int {
	list := nums
    for i, j := 0, len(list) - 1; i < j;{
        for ; list[i] % 2 == 1 && i < j; {
            i++
            
        }
        for ; list[j] % 2 == 0 && i < j; {
            j--
        }
        temp := list[i]
        list[i] = list[j]
        list[j] = temp
        i++
        j--
    }
    return list
}