/*
题目描述：  剑指 Offer 57 - II. 和为s的连续正数序列
难度：  简单
编写日期：   2022 年 2 月 15 日 晚 20：00
重写日期：  
*/
/*题目：  

	输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

    序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。


    示例 1：

    输入：target = 9
    输出：[[2,3,4],[4,5]]
    示例 2：

    输入：target = 15
    输出：[[1,2,3,4,5],[4,5,6],[7,8]]

    限制：

    1 <= target <= 10^5

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
func findContinuousSequence(target int) [][]int {
    res := make([][]int, 0)
    for i, j := 1, 2; i <= target/2; j++ {
        if sum(i, j) == target {
            temp := make([]int, j-i+1)
            for a := 0; a < j-i+1; a++ {
                temp[a] = i + a
            }
            res = append(res, temp)
        } else if sum(i, j) > target {
            i++
            j = i
        }
    }
    return res
}

func sum(i, j int) int{
    i--
    sum_i := i * (i + 1) / 2
    sum_j := j * (j + 1) / 2
    return sum_j - sum_i
}