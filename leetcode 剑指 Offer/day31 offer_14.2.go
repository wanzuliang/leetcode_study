/*
题目描述：  剑指 Offer 14- II. 剪绳子 II
难度：  中等
编写日期：   2022 年 3 月 1 日 13：00
重写日期：   
*/
/*题目：  

	给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m - 1] 。请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

    答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

    示例 1：

    输入: 2
    输出: 1
    解释: 2 = 1 + 1, 1 × 1 = 1
    示例 2:

    输入: 10
    输出: 36
    解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36

    提示：

    2 <= n <= 1000

*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
	res := nthUglyNumber(1690)
    fmt.Println( res )
}

// 默认代码模版
func cuttingRope(n int) int {
    if n < 2 {
        return 0
    }
    if n == 2 {
        return 1
    }
    if n == 3 {
        return 2
    }
    a := n / 3
    b := n % 3
    c := a / 20
    d := float64(a % 20)
    res := 0
    if b == 0 {
        res = int(math.Pow(3, d)) % 1000000007
    } else if b == 1 {
        res = int(math.Pow(3, d-1) * 4) % 1000000007
    } else {
        res = int(math.Pow(3, d) * 2) % 1000000007
    }
    temp := int( math.Pow(3, float64(20)) )
    for c > 0 {
        res = (res * temp) % 1000000007
        c--
    }
    return res
}
