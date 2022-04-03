/*
题目描述：  33. 搜索旋转排序数组
难度：  中等
编写日期：   2022 年 4 月 3 日 20：30
重写日期：   
*/
/*题目：  
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

 

示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1
 

提示：

1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-10^4 <= target <= 10^4
 

进阶：你可以设计一个时间复杂度为 O(log n) 的解决方案吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*  

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
    n := len(nums) // 数组长度
    // step1: 寻找数组最小数的下标
    l, r := 0, n-1
    for l < r {
        mid := (l+r) >> 1
        if nums[mid] > nums[r] {
            l = mid + 1     
        } else {
            r = mid
        }
    }
    smallest := r// 最小数的下标，也是整个数组的偏移量

    // step2: 尝试还原数组（用整除）
    l, r = 0, n-1
    for l <= r {
        mid := (l+r) >> 1
        realMid := (mid + smallest) % len(nums) // 未旋转之前的中间数的 “当前下标”
        if nums[realMid] == target {
            return realMid
        }
        if nums[realMid] < target {
            l = mid + 1
        } else {
            r = mid -1
        }
    }
    return -1
}
