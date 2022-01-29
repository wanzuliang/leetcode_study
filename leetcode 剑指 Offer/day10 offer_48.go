/*
题目描述：  剑指 Offer 48. 最长不含重复字符的子字符串
难度：  中等
编写日期：   2022 年 1 月 29 日 晚 21：00
重写日期：  
*/
/*题目：  
    请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

    示例 1:
    输入: "abcabcbb"
    输出: 3 
    解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

    示例 2:
    输入: "bbbbb"
    输出: 1
    解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

    示例 3:
    输入: "pwwkew"
    输出: 3

    解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
        请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
     

    提示：
    s.length <= 40000
*/
/*  
    
*/
package main

import (
    "fmt"
)

func main() {
    d := lengthOfLongestSubstring("dvdf") 
    fmt.Println(d)
    d = lengthOfLongestSubstring("aaaaab") 
    fmt.Println(d)
    d = lengthOfLongestSubstring("abbbbbb") 
    fmt.Println(d)
    d = lengthOfLongestSubstring("abcabcbb") 
    fmt.Println(d)
    d = lengthOfLongestSubstring("bbbbb") 
    fmt.Println(d)
}

// 默认代码模版
func lengthOfLongestSubstring(s string) int {
    list := [256]int{}
    max := 0
    head := 0
    tail := 0
    for i := 1; i <= len(s); i++ {
        if list[s[i-1]] <= head {
            list[s[i-1]] = 0
        }
        if list[s[i-1]] != 0 {
            head = list[s[i-1]]
            list[s[i-1]] = i
        } else {
            list[s[i-1]] = i
        }
        tail ++
        if tail-head > max {
            max = tail-head
        }
    }
    return max
}