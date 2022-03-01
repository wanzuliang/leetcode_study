/*
题目描述：  剑指 Offer 43. 1～n 整数中 1 出现的次数
难度：  困难
编写日期：   2022 年 3 月 1 日 13：00
重写日期：   
*/
/*题目：  

	输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

    例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

    示例 1：

    输入：n = 12
    输出：5
    示例 2：

    输入：n = 13
    输出：6
     

    限制：

    1 <= n < 2^31

*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    // fmt.Println( countDigitOne(9) )
    // fmt.Println( countDigitOne(99) )
    fmt.Println( countDigitOne(13) )
    fmt.Println( countDigitOne(12) )
    // fmt.Println( countDigitOne(9999) )
    // fmt.Println( countDigitOne(99999) )
}

// 默认代码模版
func countDigitOne(n int) int {
    bit := 1
    lastbit := 0
    sum := 0
    all := n
    for n > 0 {
        num := n % 10
        n = n / 10
        if num > 1 {
            sum += bit + lastbit * num
        } else if num == 1 {
            sum += all % bit + 1 + lastbit * num
        } else {
            sum +=  lastbit * num
        }
        lastbit += bit + lastbit * 9
        bit *= 10
    }
    return sum
}