/*
题目描述：  面试题32 - I. 从上到下打印二叉树
难度：  中等
编写日期：   2022 年 1 月 21 日 晚19：30
重写日期：  
*/
/*题目：  
    从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。


    例如:
    给定二叉树: [3,9,20,null,null,15,7],

        3
       / \
      9  20
        /  \
       15   7
    返回：

    [3,9,20,15,7]

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
func levelOrder(root *TreeNode) []int {
    nodeListA := make([]*TreeNode, 0)
    if root != nil {
        nodeListA = append(nodeListA, root)
    }
    list := make([]int, 0)
    for len(nodeListA) > 0 {
        nodeListB := make([]*TreeNode, 0)
        for i := 0; i < len(nodeListA); i++ {
            list = append(list, nodeListA[i].Val)
            if nodeListA[i].Left != nil {
                nodeListB = append(nodeListB, nodeListA[i].Left)
            }
            if nodeListA[i].Right != nil {
                nodeListB = append(nodeListB, nodeListA[i].Right)
            }
        }
        nodeListA = make([]*TreeNode, len(nodeListB), cap(nodeListB))
        copy(nodeListA, nodeListB)
    }
    return list
}