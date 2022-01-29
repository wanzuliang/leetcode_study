/*
题目描述：  剑指 Offer 46. 把数字翻译成字符串
难度：  中等
编写日期：   2022 年 1 月 29 日 晚 20：00
重写日期：  
*/
/*题目：  
    给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
    
    示例 1:
    输入: 12258
    输出: 5
    解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

    提示：

    0 <= num < 231
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
func translateNum(num int) int {
    str := strconv.Itoa(num)
    // str := strconv.FormatInt(num,10)
    befor_A, befor_B := 1, 1
    now := 1
    for i := 1; i < len(str); i++ {
        // d, _ := strconv.ParseInt(str[i-1:i+1], 10, 32)
        d, _ := strconv.Atoi(str[i-1:i+1])
        if  d > 9 && d < 26 {
            now = befor_A + befor_B
        } else {
            now = befor_B
        }
        befor_A = befor_B
        befor_B = now
    }
    return now
}