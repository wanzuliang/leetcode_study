/*
题目描述：  22. 括号生成
难度：  中等
编写日期：   2022 年 3 月 26 日 22：30
重写日期：   
*/
/*题目：  
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

 

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
 

提示：

1 <= n <= 8

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( generateParenthesis(2))
}

// 默认代码模版
var temp [8][]string
func generateParenthesis(n int) []string {
    if n == 1 {
        temp[0] = []string{"()"}
        return temp[0]
    }
    if len(temp[n-1]) != 0 {
        return temp[n-1]
    }
    res := make([]string, 0)
    strs := generateParenthesis(n-1)
    for _, s := range strs {
        s0 := "(" + s + ")"
        res = append(res, s0)
        s1, s2 := "()" + s, s + "()"
        if s1 == s2 {
           res = append(res, s1)
        } else {
           res = append(res, s1)
           res = append(res, s2)
        }
    }
    for n, s := range res {
        for i := n+1; i < len(res); i++ {
            if ss == res[i] {
                res[i] = res[len(res)-1]
                res = res[:len(res)-1]
                i--
            }
        }
    }
    temp[n-1] = res
    return res
}

func my(ss []string) {

}