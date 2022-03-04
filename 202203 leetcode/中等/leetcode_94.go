/*
题目描述：  94. 二叉树的中序遍历
难度：  简单
编写日期：   2022 年 3 月 3 日 22：00
重写日期：   
*/
/*题目：  

	给定一个二叉树的根节点 root ，返回它的 中序 遍历。

	 

	示例 1：


	输入：root = [1,null,2,3]
	输出：[1,3,2]
	示例 2：

	输入：root = []
	输出：[]
	示例 3：

	输入：root = [1]
	输出：[1]
	示例 4：


	输入：root = [1,2]
	输出：[2,1]
	示例 5：


	输入：root = [1,null,2]
	输出：[1,2]
	 

	提示：

	树中节点数目在范围 [0, 100] 内
	-100 <= Node.val <= 100
	 

	进阶: 递归算法很简单，你可以通过迭代算法完成吗？

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( "T" )
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
func inorderTraversal(root *TreeNode) []int {	// 	递归
    if root == nil {
        return nil
    }
    l := inorderTraversal(root.Left)
    r := inorderTraversal(root.Right)
    res := make([]int, 0)
    res = append(res, l...)
    res = append(res, root.Val)
    res = append(res, r...)
    return res
}


func inorderTraversal(root *TreeNode) []int {	//	迭代
	stack := make([]*TreeNode, 0)
    res := make([]int, 0)
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        res = append(res, root.Val)
        stack = stack[:len(stack)-1]
        root = root.Right
    }
    return res
}