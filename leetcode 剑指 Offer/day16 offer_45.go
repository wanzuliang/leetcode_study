/*
题目描述：  剑指 Offer 45. 把数组排成最小的数
难度：  中等
编写日期：   2022 年 2 月 6 日 晚 21：00
重写日期：  
*/
/*题目：  
    输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

    示例 1:

    输入: [10,2]
    输出: "102"
    示例 2:

    输入: [3,30,34,5,9]
    输出: "3033459"
     

    提示:

    0 < nums.length <= 100
    说明:

    输出结果可能非常大，所以你需要返回一个字符串而不是整数
    拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
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
func minNumber(nums []int) string {
    // nums[0] = 3
    // nums[1] = 30
    // if IsWeightA(nums[0], nums[1]) {
    //     return "0 max"
    // } else {
    //     return "1 max"
    // }
    result := make([]byte,0)
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if IsWeightA(nums[j],nums[i]) {
                temp := nums[j]
                nums[j] = nums[i]
                nums[i] = temp
            }
        }
        temp := strconv.Itoa(nums[i])
        for j := 0; j < len(temp); j++ {
            result = append(result, temp[j] )
        }
    }
    return string(result)
}

func IsWeightA(a, b int) bool {
    sab, _ := strconv.Atoi(strconv.Itoa( a ) + strconv.Itoa( b ))
    sba, _ := strconv.Atoi(strconv.Itoa( b ) + strconv.Itoa( a ))
    if sab < sba {
        return true
    }
    return false
}