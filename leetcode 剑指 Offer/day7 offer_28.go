/*
题目描述：  剑指 Offer 28. 对称的二叉树
难度：  简单
编写日期：   2022 年 1 月 24 日 晚19：50
重写日期：  
*/
/*题目：  
    请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

    例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

        1
       / \
      2   2
     / \ / \
    3  4 4  3

    但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

        1
       / \
      2   2
       \   \
       3    3

    示例 1：

    输入：root = [1,2,2,3,4,4,3]
    输出：true
    示例 2：

    输入：root = [1,2,2,null,3,null,3]
    输出：false

    限制：

    0 <= 节点个数 <= 1000
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isSame(root.Left, root.Right)
}

func isSame(LN *TreeNode, RN *TreeNode) bool {
    if LN == nil && RN == nil{
        return true
    }
    if LN == nil && RN != nil {
        return false
    }
    if LN != nil && RN == nil {
        return false
    }
    if LN.Val == RN.Val {
        return isSame(LN.Left, RN.Right) && isSame(LN.Right, RN.Left)
    }
    return false
}