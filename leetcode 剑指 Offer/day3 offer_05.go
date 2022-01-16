/*
题目描述：  剑指 Offer 05. 替换空格
难度：  简单
编写日期：  之前写过 2021/12/03 17:00
重写日期：   2022 年 1 月 16 日 晚20：00
*/
/*题目：  
    请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

    示例 1：

    输入：s = "We are happy."
    输出："We%20are%20happy."

    限制：

    0 <= s 的长度 <= 10000

*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println("T")
}

// 默认代码模版
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func replaceSpace(s string) string {
    length := len(s)
    newLen := length
    for _, v := range s {
        if v == ' ' {
            newLen += 2
        }
    }
    newStr := make([]byte, newLen)
    for i, j := length-1, newLen-1; i >= 0 ; i-- {
        if s[i] == ' ' {
            newStr[j] = '0'
            newStr[j-1] = '2'
            newStr[j-2] = '%'
            j = j-3
        } else {
            newStr[j] = s[i]
            j--
        }
    }
    return string(newStr)
}