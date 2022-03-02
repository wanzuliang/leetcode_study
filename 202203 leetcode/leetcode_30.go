/*
题目描述：  35. 搜索插入位置
难度：  简单
编写日期：   2022 年 3 月 2 日 22：00
重写日期：   
*/
/*题目：  

	给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

    请必须使用时间复杂度为 O(log n) 的算法。

     

    示例 1:

    输入: nums = [1,3,5,6], target = 5
    输出: 2
    示例 2:

    输入: nums = [1,3,5,6], target = 2
    输出: 1
    示例 3:

    输入: nums = [1,3,5,6], target = 7
    输出: 4
    示例 4:

    输入: nums = [1,3,5,6], target = 0
    输出: 0
    示例 5:

    输入: nums = [1], target = 0
    输出: 0
     

    提示:

    1 <= nums.length <= 104
    -104 <= nums[i] <= 104
    nums 为无重复元素的升序排列数组
    -104 <= target <= 104

    来源：力扣（LeetCode）
    链接：https://leetcode-cn.com/problems/search-insert-position
    著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( "T" )
}

// 默认代码模版
func searchInsert(nums []int, target int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] >= target {
            return i
        } else {
            res++
        }
    }
    return res
}