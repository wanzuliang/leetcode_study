/*
题目描述：  1. 两数之和
难度：  简单
编写日期：   2022 年 3 月 7 日 22：20
重写日期：   
*/
/*题目：  

	给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

    你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

    你可以按任意顺序返回答案。

     

    示例 1：

    输入：nums = [2,7,11,15], target = 9
    输出：[0,1]
    解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
    示例 2：

    输入：nums = [3,2,4], target = 6
    输出：[1,2]
    示例 3：

    输入：nums = [3,3], target = 6
    输出：[0,1]
     

    提示：

    2 <= nums.length <= 104
    -109 <= nums[i] <= 109
    -109 <= target <= 109
    只会存在一个有效答案
    进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？

    来源：力扣（LeetCode）
    链接：https://leetcode-cn.com/problems/two-sum
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
func twoSum(nums []int, target int) []int {
    hashTable := map[int]int{}
    for i, x := range nums {
        if p, ok := hashTable[target-x]; ok {
            return []int{p, i}
        }
        hashTable[x] = i
    }
    return nil
}




/////
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}


///////
func twoSum(nums []int, target int) []int {
    temp := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        temp[i] = i
    }
    quickSort(nums, 0, len(nums)-1, temp)
    for i, j := 0, len(nums)-1; i < j; {
        sum := nums[i] + nums[j]
        if sum == target {
            return []int{temp[i], temp[j]}
        }
        if sum < target {
            i++
        } else {
            j--
        }
    }
    return nil
}

func quickSort(arr []int, left, right int, index []int) {
    if left >= right {
        return
    }
    cur, lo := left+1, left
    for cur <= right {
        if arr[cur] <= arr[left] {
            arr[lo+1], arr[cur] = arr[cur], arr[lo+1]
            index[lo+1], index[cur] = index[cur], index[lo+1]
            lo++
        }
        cur++
    }
    arr[left], arr[lo] = arr[lo], arr[left]
    index[left], index[lo] = index[lo], index[left]
    quickSort(arr, left, lo-1, index)
    quickSort(arr, lo+1, right, index)
}
        

///////////////////////////
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            for j in range(i+1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j] 
        return []
    # def twoSum(self, nums: List[int], target: int) -> List[int]:
    #     hashtable = dict()
    #     for i, num in enumerate(nums):
    #         if target - num in hashtable:
    #             return [hashtable[target - num], i]
    #         hashtable[nums[i]] = i
    #     return []