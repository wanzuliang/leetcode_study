/*
题目描述：  剑指 Offer 32 - II. 从上到下打印二叉树 II
难度：  简单
编写日期：   2022 年 1 月 21 日 晚20：30
重写日期：  
*/
/*题目：  
    从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

    例如:
    给定二叉树: [3,9,20,null,null,15,7],
        3
       / \
      9  20
        /  \
       15   7
    返回其层次遍历结果：
    [
      [3],
      [9,20],
      [15,7]
    ]

    提示：
    节点总数 <= 1000
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
func levelOrder(root *TreeNode) [][]int {
    nodeListA := make([]*TreeNode, 0)
    if root != nil {
        nodeListA = append(nodeListA, root)
    }
    list := make([][]int, 0)
    for len(nodeListA) > 0 {
        nodeListB := make([]*TreeNode, 0)
        line := make([]int, 0)
        for i := 0; i < len(nodeListA); i++ {
            line = append(line, nodeListA[i].Val)
            if nodeListA[i].Left != nil {
                nodeListB = append(nodeListB, nodeListA[i].Left)
            }
            if nodeListA[i].Right != nil {
                nodeListB = append(nodeListB, nodeListA[i].Right)
            }
        }
        list = append(list, line)
        nodeListA = make([]*TreeNode, len(nodeListB), cap(nodeListB))
        copy(nodeListA, nodeListB)
    }
    return list
}