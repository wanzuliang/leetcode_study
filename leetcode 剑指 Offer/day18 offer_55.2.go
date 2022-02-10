/*
题目描述：  剑指 Offer 55 - II. 平衡二叉树
难度：  简单
编写日期：   2022 年 2 月 10 日 晚 20：00
重写日期：  
*/
/*题目：  
    输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

	示例 1:

	给定二叉树 [3,9,20,null,null,15,7]

	    3
	   / \
	  9  20
	    /  \
	   15   7
	返回 true 。

	示例 2:

	给定二叉树 [1,2,2,3,3,null,null,4,4]

	       1
	      / \
	     2   2
	    / \
	   3   3
	  / \
	 4   4
	返回 false 。

	限制：

	0 <= 树的结点个数 <= 10000


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
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    le := maxDepth(root.Left)
	ri := maxDepth(root.Right)
    if le - ri > 1 || le - ri < -1 {
        return false
    }
    return isBalanced(root.Left) && isBalanced(root.Right)
}

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