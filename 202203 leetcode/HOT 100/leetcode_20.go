/*
题目描述：  20. 有效的括号
难度：  简单
编写日期：   2022 年 3 月 25 日 21：30
重写日期：   
*/
/*题目：  
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
 

示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
示例 4：

输入：s = "([)]"
输出：false
示例 5：

输入：s = "{[]}"
输出：true
 

提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    str := "({))"
    fmt.Println(str, isValid(str))
}

// 默认代码模版
func isValid(str string) bool {
    if len(str) == 0 {
        return true
    }
    stack := make([]byte, 0)
    stack = append(stack, '0')
    //(){}[]
    for i := 0; i < len(str); i++ {
        if str[i] == '(' || str[i] == '[' || str[i] == '{' {
            stack = append(stack, str[i])
            continue
        }
        end := len(stack)-1
        switch str[i] {
        case ')':
            if stack[end] == '(' {
                stack = stack[:end]
                break
            } else {
                return false
            }
        case ']':
            if stack[end] == '[' {
                stack = stack[:end]
                break
            } else {
                return false
            }
        case '}':
            if stack[end] == '{' {
                stack = stack[:end]
                break
            } else {
                return false
            }
        default:
            return false
        }
    }
    if len(stack) == 1 {
        return true
    } else {
        return false
    }
}