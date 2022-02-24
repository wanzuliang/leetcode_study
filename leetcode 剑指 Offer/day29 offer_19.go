/*
题目描述：  面试题19. 正则表达式匹配
难度：  困难
编写日期：   2022 年 2 月 24 日 晚 17：00
重写日期：  
*/
/*题目：  

	请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

    示例 1:

    输入:
    s = "aa"
    p = "a"
    输出: false
    解释: "a" 无法匹配 "aa" 整个字符串。
    示例 2:

    输入:
    s = "aa"
    p = "a*"
    输出: true
    解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
    示例 3:

    输入:
    s = "ab"
    p = ".*"
    输出: true
    解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
    示例 4:

    输入:
    s = "aab"
    p = "c*a*b"
    输出: true
    解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
    示例 5:

    输入:
    s = "mississippi"
    p = "mis*is*p*."
    输出: false
    s 可能为空，且只包含从 a-z 的小写字母。
    p 可能为空，且只包含从 a-z 的小写字母以及字符 . 和 *，无连续的 '*'。

*/
/*  
    
*/
package main

import (
    "fmt"
)

func main() {
    s := "aab"
    p := "c*a*b"        //true
    fmt.Println( isMatch(s, p))

    s = "aa"
    p = "a"             //false
    fmt.Println( isMatch(s, p))

    s = "mississippi"
    p = "mis*is*p*."    //false
    fmt.Println( isMatch(s, p))
    s = "ab"
    p = ".*"            //true
    fmt.Println( isMatch(s, p))
    s = "aa"
    p = "a*"            //true
    fmt.Println( isMatch(s, p))
}

// 默认代码模版
func isMatch(s string, p string) bool {
    if p == ""{
        return s == ""
    }
    old := make([]bool, len(s)+1)
    current := make([]bool, len(s)+1)
    i := 0
    for i < len(p) {
        if i+1 < len(p) && p[i+1] == '*' {
            // 读到x*
            if i==0{
                // x*出现在开头
                current[0] = true
                for j := 1; j < len(s)+1; j++ {
                    if s[j-1] == p[i] || p[i] == '.' {
                        current[j] = true
                    } else {
                        break
                    }
                }
            } else {
                // x*前面还有
                for j := 0; j < len(s)+1 ; j++ {
                    if old[j] {
                        current[j] = true
                        for k := j; k < len(s); k++ {
                            if s[k] == p[i] || p[i] == '.' {
                                current[k+1] = true
                            }else{
                                break
                            }
                        }
                    }
                }
            }           
            i = i + 2
        } else if p[i] == '.' {
            if i == 0 {
                if len(s) > 0 {
                    current[1] = true
                } else {
                    return false
                }
            } else {
                for j := 0; j < len(s)+1; j++ {
                    if old[j] {
                        if j+1 < len(s)+1 {
                            current[j+1] = true
                        }
                    }
                }
            }
            i++
        } else {
            if i == 0 {
                if len(s) > 0 && s[0] == p[0] {
                    current[1] = true
                }else{
                    return false
                }
            } else {
                for j := 0; j < len(s)+1; j++ {
                    if old[j]{
                        if j+1 < len(s)+1 && s[j] == p[i]{
                            current[j+1] = true
                        }
                    }
                }
            }
            i++
        }
        //fmt.Println(current)
        old = current
        current = make([]bool, len(s)+1)
    }
    return old[len(s)]
}