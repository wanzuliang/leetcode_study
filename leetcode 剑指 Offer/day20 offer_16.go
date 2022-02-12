/*
题目描述：  剑指 Offer 16. 数值的整数次方
难度：  中等
编写日期：   2022 年 2 月 12 日 晚 20：00
重写日期：  
*/
/*题目：  

	实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。

    示例 1：

    输入：x = 2.00000, n = 10
    输出：1024.00000
    示例 2：

    输入：x = 2.10000, n = 3
    输出：9.26100
    示例 3：

    输入：x = 2.00000, n = -2
    输出：0.25000
    解释：2-2 = 1/22 = 1/4 = 0.25
     
    提示：

    -100.0 < x < 100.0
    -231 <= n <= 231-1
    -104 <= xn <= 104

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
func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    } else if n < 0 {
        x = 1 / x
        n = -n
    }
    res := x
    if n % 2 == 0 {
        res = myPow(x, n / 2)
        res = res * res
    } else {
        res = myPow(x, n - 1) * x
    }
    return res
}