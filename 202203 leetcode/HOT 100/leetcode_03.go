/*
题目描述：  3. 无重复字符的最长子串
难度：  中等
编写日期：   2022 年 3 月 6 日 22：16
重写日期：   
*/
/*题目：  

	给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

     

    示例 1:

    输入: s = "abcabcbb"
    输出: 3 
    解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
    示例 2:

    输入: s = "bbbbb"
    输出: 1
    解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
    示例 3:

    输入: s = "pwwkew"
    输出: 3
    解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
         请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
     

    提示：

    0 <= s.length <= 5 * 104
    s 由英文字母、数字、符号和空格组成

    来源：力扣（LeetCode）
    链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
    著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( "T" )
}

// 默认代码模版
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    temp := make([]int, 256)
    res := 0
    long := 0
    for i, j := 0, 0; j < len(s); j++ {
        if temp[s[j]] > 0 {
            i = temp[s[j]]
            j = i
            temp = make([]int, 256)
            long = 0
        }
        temp[s[j]] = j + 1
        long++
        if long > res {
            res = long
        }
    }
    return res
}

// 官方方法
func lengthOfLongestSubstring(s string) int {
    // 哈希集合，记录每个字符是否出现过
    m := map[byte]int{}
    n := len(s)
    // 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
    rk, ans := -1, 0
    for i := 0; i < n; i++ {
        if i != 0 {
            // 左指针向右移动一格，移除一个字符
            delete(m, s[i-1])
        }
        for rk + 1 < n && m[s[rk+1]] == 0 {
            // 不断地移动右指针
            m[s[rk+1]]++
            rk++
        }
        // 第 i 到 rk 个字符是一个极长的无重复字符子串
        ans = max(ans, rk - i + 1)
    }
    return ans
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}
