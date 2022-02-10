/*
题目描述：  剑指 Offer 55 - I. 二叉树的深度
难度：  简单
编写日期：   2022 年 2 月 10 日 晚 20：00
重写日期：  
*/
/*题目：  
    输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

	例如：

	给定二叉树 [3,9,20,null,null,15,7]，

	    3
	   / \
	  9  20
	    /  \
	   15   7
	返回它的最大深度 3 。

	提示：

	节点总数 <= 10000

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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	le := maxDepth(root.Left) + 1
	ri := maxDepth(root.Right) + 1
	if ri > le {
		return ri
	} else {
		return le
	}
}