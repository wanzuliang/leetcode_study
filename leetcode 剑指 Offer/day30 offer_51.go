/*
题目描述：  剑指 Offer 51. 数组中的逆序对
难度：  困难
编写日期：   2022 年 2 月 28 日 晚 20：00
重写日期：   
*/
/*题目：  

	在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

    示例 1:

    输入: [7,5,6,4]
    输出: 5

    限制：

    0 <= 数组长度 <= 50000
*/
/*  
    没写
*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( "T" )
}

// 默认代码模版
func reversePairs(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    tmp := make([]int, len(nums))
    return merge(nums, tmp, 0, len(nums)-1)
}

func merge(nums, tmp []int, start, end int) int {
    if start >= end {
        return 0
    }
    mid := (start + end) / 2
    leftCnt := merge(nums, tmp, start, mid)
    rightCnt := merge(nums, tmp, mid+1, end)
    mergeCnt := 0
    index := start
    i, j := start, mid+1
    for i <= mid && j <= end {
        if nums[i] > nums[j] {
            mergeCnt = mergeCnt + mid - i + 1
            tmp[index] = nums[j]
            j++
        } else if nums[i] <= nums[j] {
            tmp[index] = nums[i]
            i++
        }
        index++
    }
    for i <= mid {
        tmp[index] = nums[i]
        i++
        index++
    }
    for j <= end {
        tmp[index] = nums[j]
        j++
        index++
    }
    for m := start; m <= end; m++ {
        nums[m] = tmp[m]
    }
    return leftCnt + rightCnt + mergeCnt

}

///////////////////////////////////////
func reversePairs(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    return merge(nums, 0, len(nums)-1)
}

func merge(nums []int, start, end int) int {
    if start >= end {
        return 0
    }
    mid := start + (end - start)/2
    cnt := merge(nums, start, mid) + merge(nums, mid + 1, end)
    tmp := []int{}
    i, j := start, mid + 1
    for i <= mid && j <= end {
        if nums[i] <= nums[j] {
            tmp = append(tmp, nums[i])
            cnt += j - (mid + 1)
            i++
        } else {
            tmp = append(tmp, nums[j])
            j++
        }
    }
    for ; i <= mid; i++ {
        tmp = append(tmp, nums[i])
        cnt += end - (mid + 1) + 1
    }
    for ; j <= end; j++ {
        tmp = append(tmp, nums[j])
    }
    for i := start; i <= end; i++ {
        nums[i] = tmp[i - start]
    }
    return cnt
}
