/*
题目描述：  53. 最大子数组和
难度：  简单
编写日期：   2022 年 3 月 2 日 22：15
重写日期：   2022 年 3 月 3 日 10：15
*/
/*题目：  

	给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

	子数组 是数组中的一个连续部分。

	 

	示例 1：

	输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
	输出：6
	解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
	示例 2：

	输入：nums = [1]
	输出：1
	示例 3：

	输入：nums = [5,4,-1,7,8]
	输出：23
	 

	提示：

	1 <= nums.length <= 105
	-104 <= nums[i] <= 104

	进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/maximum-subarray
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
func maxSubArray(nums []int) int {
    max := nums[0]
    now := nums[0]
    
    for b := 1; b < len(nums); b++ {
        if now < 0 {
            now = nums[b]
        } else {
            now += nums[b]
        }
        if now > max {
            max = now
        }
    }
    return max
}

//
func maxSubArray(nums []int) int {		//分治法
    _, max, _, _ := div(nums)
    return max
}


func div(nums []int) (maxL, max, maxR, all int) { 
    if len(nums) == 0 {
        return
    }
    if len(nums) == 1 {
        max = nums[0]
        maxL = nums[0]
        maxR = nums[0]
        all = nums[0]
        return 
    }
    LmaxL, Lmax, LmaxR, Lall := div(nums[:len(nums)/2])
    RmaxL, Rmax, RmaxR, Rall := div(nums[len(nums)/2:])
    LR_RL := LmaxR + RmaxL
    all = Lall + Rall
    //
    if Lall + RmaxL >= LmaxL {  // left max
        maxL = Lall + RmaxL
    } else if all >= LmaxL {
        maxL = all
    } else {
        maxL = LmaxL
    }
    if maxL >= Lmax {
        Lmax = maxL
    }
    //
    if LmaxR + Rall >= RmaxR {  // right max
        maxR = LmaxR + Rall
    } else if all >= RmaxR {
        maxR = all
    } else {
        maxR = RmaxR
    }
    if maxR >= Rmax {
        Rmax = maxR
    }
    // max
    if Rmax > Lmax {
        max = Rmax
    } else {
        max = Lmax
    }
    if LR_RL > max {
        max = LR_RL
    }
    if all >= max {
        max = all
        maxL = all
        maxR = all
    }
    return
}







