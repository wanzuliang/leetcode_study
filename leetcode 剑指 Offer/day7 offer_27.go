/*
题目描述：  剑指 Offer 27. 二叉树的镜像
难度：  简单
编写日期：   2022 年 1 月 24 日 晚19：30
重写日期：  
*/
/*题目：  
    请完成一个函数，输入一个二叉树，该函数输出它的镜像。

    例如输入：
            4
          /   \
         2     7
        / \   / \
       1   3 6   9

    镜像输出：
            4
          /   \
         7     2
        / \   / \
       9   6 3   1

    示例 1：

    输入：root = [4,2,7,1,3,6,9]
    输出：[4,7,2,9,6,3,1]

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
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Left != nil || root.Right != nil {
        node := root.Left
        root.Left = root.Right
        root.Right = node
    }
    if root.Left != nil {
        mirrorTree(root.Left)
    }
    if root.Right != nil {
        mirrorTree(root.Right)
    }
    return root
}