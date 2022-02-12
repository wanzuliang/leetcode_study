/*
题目描述：  剑指 Offer 07. 重建二叉树
难度：  中等
编写日期：   2022 年 2 月 12 日 晚 20：00
重写日期：  
*/
/*题目：  

	输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。

	假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

	示例 1:

	Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	Output: [3,9,20,null,null,15,7]
	示例 2:

	Input: preorder = [-1], inorder = [-1]
	Output: [-1]

	限制：

	0 <= 节点个数 <= 5000

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
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    var head TreeNode
    head.Val = preorder[0]
    head_location := 0
    for i := 0; i < len(inorder); i++ {
        if inorder[i] == preorder[0] {
            head_location = i
        }
    }
    head.Left = buildTree(preorder[1:head_location + 1], inorder[0:head_location])
    head.Right = buildTree(preorder[head_location + 1:len(preorder)], inorder[head_location + 1:len(inorder)])
    return &head
}