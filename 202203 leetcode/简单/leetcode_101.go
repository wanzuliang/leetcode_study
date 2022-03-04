/*
题目描述：  101. 对称二叉树
难度：  简单
编写日期：   2022 年 3 月 3 日 22：28
重写日期：   
*/
/*题目：  

	给你一个二叉树的根节点 root ， 检查它是否轴对称。


	示例 1：


	输入：root = [1,2,2,3,4,4,3]
	输出：true
	示例 2：


	输入：root = [1,2,2,null,3,null,3]
	输出：false

	提示：

	树中节点数目在范围 [1, 1000] 内
	-100 <= Node.val <= 100

	进阶：你可以运用递归和迭代两种方法解决这个问题吗？

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/symmetric-tree
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
func isSymmetric(root *TreeNode) bool {
    return isSameNode(root.Left, root.Right)
}

func isSameNode(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    } else if left == nil {
        return false
    } else if right == nil {
        return false
    }
    if left.Val != right.Val {
        return false
    } else {
        return isSameNode(left.Left, right.Right) && isSameNode(left.Right, right.Left) 
    }
}