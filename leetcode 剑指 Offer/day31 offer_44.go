/*
题目描述：  剑指 Offer 44. 数字序列中某一位的数字
难度：  困难
编写日期：   2022 年 3 月 1 日 13：00
重写日期：   
*/
/*题目：  

	数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

    请写一个函数，求任意第n位对应的数字。

    示例 1：

    输入：n = 3
    输出：3
    示例 2：

    输入：n = 11
    输出：0

    限制：

    0 <= n < 2^31

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
func findNthDigit(n int) int {
    num := 1
    bit := 0
    for { //    判断第n位的哪个数是几位数 比如是3位数
        temp := 9 * int(math.Pow(10, float64(bit))) * (bit + 1)
        if num + temp > n {
            break
        }
        num += temp
        bit++
    }
    ans := n - num      //  第n位对应的数字 N 对于 100(如果是3位数)的位移
    a := ans / (bit+1)  //  100后的第几个数字
    b := ans % (bit+1)  //  在这个数字 N 里的位置 m
    num = a + int(math.Pow10(bit))      // 这就是第n位的那个完整的数 N
    ans = num / int(math.Pow10(bit-b)) % 10 // 得到 N 中的位置 m 的值
    return ans
}