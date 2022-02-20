/*
题目描述：  剑指 Offer 38. 字符串的排列
难度：  中等
编写日期：   2022 年 2 月 20 日 晚 20：00
重写日期：  
*/
/*题目：  

    输入一个字符串，打印出该字符串中字符的所有排列。

    你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。


    示例:

    输入：s = "abc"
    输出：["abc","acb","bac","bca","cab","cba"]

    限制：

    1 <= s 的长度 <= 8

*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    s := "abc"
    i := 0
    fmt.Println( s[:i] + s[i+1:] )
    i = 1
    fmt.Println( s[:i] + s[i+1:] )
    // i := 2
    fmt.Println( permutation(s) )
}

// 默认代码模版
func permutation(s string) []string {
    if len(s) == 0 {
        return nil
    }
    list := make([]string, 0)
    if len(s) == 1 {
        list = append(list, string(s[0]))
        return list
    }
    for i := 0; i < len(s); i++ {
        ch := s[i]
        flag := false
        for j := 0; j < i; j++ {
            if ch == s[j] {
                flag = true
            }
        }
        if flag {
            continue
        }
        strs := permutation( s[:i] + s[i+1:] )

        for _, str := range strs {
            list = append(list, string(ch) + str)
        }
    }
    return list
}