/*
题目描述：  15. 三数之和
难度：  中等
编写日期：   2022 年 3 月 15 日 22：30
重写日期：   
*/
/*题目：  
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

 

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]
 

提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
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
func threeSum(nums []int) [][]int {                 //时间超过，
    if len(nums) < 3 {
        return [][]int{}
    }
    res := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums) && nums[i] <= 0; i++ {
        for j := i+1; j < len(nums); j++ {
            for k := len(nums)-1; k > j; k-- {
                                    
            flag := 0
            if len(res) > 0 {
                for k := len(res)-1; k >= 0 && nums[i] == res[k][0]; k-- {    //去重
                    if nums[j] == res[k][1] {
                        flag = 1
                        break
                    }
                }
            }
            if flag == 1 {
                continue
            }
                if nums[i] + nums[j] < - nums[k]{
                    break
                }
                if nums[i] + nums[j] == - nums[k] {
                    res = append(res, []int{nums[i], nums[j], nums[k]})
                    break
                }
            }
        }
    }
    return res
}

// 官方给的
func threeSum(nums []int) [][]int {
    n := len(nums)
    sort.Ints(nums)
    ans := make([][]int, 0)
 
    // 枚举 a
    for first := 0; first < n; first++ {
        // 需要和上一次枚举的数不相同
        if first > 0 && nums[first] == nums[first - 1] {
            continue
        }
        // c 对应的指针初始指向数组的最右端
        third := n - 1
        target := -1 * nums[first]
        // 枚举 b
        for second := first + 1; second < n; second++ {
            // 需要和上一次枚举的数不相同
            if second > first + 1 && nums[second] == nums[second - 1] {
                continue
            }
            // 需要保证 b 的指针在 c 的指针的左侧
            for second < third && nums[second] + nums[third] > target {
                third--
            }
            // 如果指针重合，随着 b 后续的增加
            // 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
            if second == third {
                break
            }
            if nums[second] + nums[third] == target {
                ans = append(ans, []int{nums[first], nums[second], nums[third]})
            }
        }
    }
    return ans
}
