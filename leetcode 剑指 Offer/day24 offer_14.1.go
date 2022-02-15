/*
题目描述：  剑指 Offer 14- I. 剪绳子
难度：  中等
编写日期：   2022 年 2 月 15 日 晚 20：00
重写日期：  
*/
/*题目：  

	给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

    示例 1：

    输入: 2
    输出: 1
    解释: 2 = 1 + 1, 1 × 1 = 1
    示例2:

    输入: 10
    输出: 36
    解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
    提示：

    2 <= n <= 58
*/
/*  
    考虑到只有裁成最小为 3 的小段，才能使所有段乘积最大，根据余下长度不足 3 的段化为长度为 2 的段段
*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( "T")
}

// 默认代码模版
func cuttingRope(n int) int {
    if n < 2 {
        return 0
    }else if n == 2 {
        return 1
    }else if n == 3 {
        return 2
    }else if n % 3 == 0 {
        return int(math.Pow(3, float64(n/3)))
    } else if n % 3 == 1 {
        return int(math.Pow(3, float64(n/3) - 1)) * 2 * 2
    } else if n % 3 == 2 {
        return int(math.Pow(3, float64(n/3))) * 2
    }
    return -1
}