/*
题目描述：  剑指 Offer 50. 第一个只出现一次的字符
难度：  简单
编写日期：   2022 年 1 月 20 日 晚21：30
重写日期：  
*/
/*题目：  
    在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

    示例 1:
    输入：s = "abaccdeff"
    输出：'b'
    示例 2:

    输入：s = "" 
    输出：' '
    
    限制：
    0 <= s 的长度 <= 50000
*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    // sum := make([]int, 1)
    // sum := []int{0,1,2,3,4}
    // i := ' '//32
    // i := 'a'//97
    str := "itwqbtcdprfsuprkrjkausiterybzncbmdvkgljxuekizvaivszowqtmrttiihervpncztuoljftlxybpgwnjb"
    ch := firstUniqChar(str)
    fmt.Println( "'" + string(ch) + "'")
}

// 默认代码模版
func firstUniqChar(s string) byte {
    if len(s) == 0{
        return ' '
    }
    sum := [26]int{}
    rank := 1
    for i := 0; i < len(s); i++ {
        ch := s[i]
        if sum[ch - 'a'] == 0 {
            sum[ch - 'a'] = rank
            rank++
        } else {
            sum[ch - 'a'] = -1
        }
    }
    min := 100
    ch := byte(' ')
    for i := 0; i < 26; i++ {
        fmt.Println(    sum[i])
        if sum[i] < min && sum[i] > 0 {
            min = sum[i]
            ch = byte(i + 'a')
        }
    }
    return ch
}