/*
题目描述：  剑指 Offer 33. 二叉搜索树的后序遍历序列
难度：  中等
编写日期：   2022 年 2 月 12 日 晚 20：00
重写日期：  
*/
/*题目：  

	输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。


    参考以下这颗二叉搜索树：

         5
        / \
       2   6
      / \
     1   3
    示例 1：

    输入: [1,6,3,2,5]
    输出: false
    示例 2：

    输入: [1,3,2,6,5]
    输出: true

    提示：

    数组长度 <= 1000

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
func verifyPostorder(postorder []int) bool {
    if len(postorder) <= 1 {
        return true
    }
    end := len(postorder)-1
    mid := end
    for i := 0; i < end; i++ {
        if postorder[i] > postorder[end] {
            mid = i
            break
        }
    }
    for i := mid; i < end; i++ {
        if postorder[i] < postorder[end] {
            return false
        }
    }
    return verifyPostorder(postorder[0:mid]) && verifyPostorder(postorder[mid:end])
}