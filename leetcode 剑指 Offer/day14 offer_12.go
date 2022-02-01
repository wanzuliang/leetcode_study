/*
题目描述：  剑指 Offer 12. 矩阵中的路径
难度：  中等
编写日期：   2022 年 2 月 1 日 晚 19：00
重写日期：  
*/
/*题目：  
    给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

    单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

    例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。
    [A] [B] [C] E
    S   F   [C] S
    A   [D] [E] E

    示例 1：

    输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
    输出：true
    示例 2：

    输入：board = [["a","b"],["c","d"]], word = "abcd"
    输出：false
     

    提示：

    1 <= board.length <= 200
    1 <= board[i].length <= 200
    board 和 word 仅由大小写英文字母组成
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
func exist(board [][]byte, word string) bool {
    if board == nil || len(board) == 0 || len(board[0]) == 0 {
        return false
    }
    visited := make([][]bool, len(board))
    for i := 0; i < len(board); i++ {
        visited[i] = make([]bool, len(board[0]))
    }
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if dfs(board, word, visited, i, j, 0) {
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, word string, visited [][]bool, i, j, start int) bool {
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || word[start] != board[i][j] || visited[i][j] {
            return false
    }
    if start == len(word) - 1 {
        return true
    }
    visited[i][j] = true
    flag := false
    flag = dfs(board, word, visited, i - 1, j, start + 1) ||  dfs(board, word, visited, i, j + 1, start + 1) ||  dfs(board, word, visited, i + 1, j, start + 1) ||  dfs(board, word, visited, i, j - 1, start + 1)
    visited[i][j] = false
    return flag    
}