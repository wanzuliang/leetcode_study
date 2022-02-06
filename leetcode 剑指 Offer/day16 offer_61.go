/*
题目描述：  剑指 Offer 61. 扑克牌中的顺子
难度：  简单
编写日期：   2022 年 2 月 6 日 晚 21：00
重写日期：  
*/
/*题目：  
    从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

    示例 1:

    输入: [1,2,3,4,5]
    输出: True
     

    示例 2:

    输入: [0,0,1,2,5]
    输出: True
     

    限制：

    数组长度为 5 

    数组的数取值为 [0, 13] .
*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    input := []int{8,7,11,0,9}
    fmt.Println( isStraight(input) )
}

// 默认代码模版
func isStraight(nums []int) bool {
    wang := 0
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[j] < nums[i] {
                temp := nums[i]
                nums[i] = nums[j]
                nums[j] = temp
            }
        }
        if nums[i] == 0 {
            wang++
        }
    }
    next := nums[0 + wang]
    for i := 0 + wang; i < len(nums); {
        if nums[i] == next {
            next ++
            i ++
        } else {
            if wang > 0 {
                next ++
                wang --
            } else {
                return false
            }
        }
    }
    return true
}