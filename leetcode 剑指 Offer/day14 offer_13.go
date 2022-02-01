/*
题目描述：  剑指 Offer 13. 机器人的运动范围
难度：  中等
编写日期：   2022 年 2 月 1 日 晚 20：00
重写日期：  
*/
/*题目：  
    地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

    示例 1：

    输入：m = 2, n = 3, k = 1
    输出：3
    示例 2：

    输入：m = 3, n = 1, k = 0
    输出：1
    提示：

    1 <= n,m <= 100
    0 <= k <= 20
*/
/*  

*/
package main

import (
    "fmt"
)

func main() {

    fmt.Println( movingCount(2,3,1))
}

// 默认代码模版
func movingCount(m int, n int, k int) int {
    visited := make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }

    bfs(visited, m, n, 0, 0, k)
    
    sum := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if visited[i][j] {
                sum++
            }
        }
    }
    return sum
}

func bfs(visited [][]bool, m, n, i, j, k int ) {
    if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] == true {
        return
    }
    if i % 10 + i / 10 + j % 10 + j / 10 <= k {
        visited[i][j] = true
        bfs(visited, m, n, i+1, j, k)
        bfs(visited, m, n, i-1, j, k)
        bfs(visited, m, n, i, j+1, k)
        bfs(visited, m, n, i, j-1, k)
    } else {
        visited[i][j] = false
    }
}