/*
题目描述：  剑指 Offer 32 - III. 从上到下打印二叉树 III
难度：  简单
编写日期：   2022 年 1 月 21 日 晚21：00
重写日期：  
*/
/*题目：  
    请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

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
      [20,9],
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
    layer := 1
    for len(nodeListA) > 0 {
        nodeListB := make([]*TreeNode, 0)
        line := make([]int, len(nodeListA))
        for i := 0; i < len(nodeListA); i++ {
            if layer % 2 != 0 {
                line[i] = nodeListA[i].Val
            } else {
                line[len(nodeListA)-1 - i] = nodeListA[i].Val
            }
            if nodeListA[i].Left != nil {
                nodeListB = append(nodeListB, nodeListA[i].Left)
            }
            if nodeListA[i].Right != nil {
                nodeListB = append(nodeListB, nodeListA[i].Right)
            }
        }
        layer++
        list = append(list, line)
        nodeListA = make([]*TreeNode, len(nodeListB), cap(nodeListB))
        copy(nodeListA, nodeListB)
    }
    return list
}