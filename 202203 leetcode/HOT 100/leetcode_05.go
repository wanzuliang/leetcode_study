/*
题目描述：  5. 最长回文子串
难度：  中等
编写日期：   2022 年 3 月 12 日 23：00
重写日期：   
*/
/*题目：  
给你一个字符串 s，找到 s 中最长的回文子串。

 

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
 

提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
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
func longestPalindrome(s string) string {
    max := 1
    l, r := 0, 0
    for n := 0; n < len(s); n++ {
        for i := 1; 0 <= n-i && n+i < len(s); i++ { //奇数回文
            if s[n-i] == s[n+i] {
                if i*2+1 > max {
                    max = i*2 + 1
                    l = n - i
                    r = n + i
                }
            } else {
                break
            }
        }
        for i := 1; 0 <= n-i+1 && n+i < len(s); i++ { //偶数回文
            if s[n-i+1] == s[n+i] {
                if 2*i > max {
                    max = 2 * i
                    l = n - i + 1
                    r = n + i
                }
            } else {
                break
            }
        }
    }
    return s[l:r+1]
}