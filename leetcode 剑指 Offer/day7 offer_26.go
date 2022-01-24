/*
题目描述：  剑指 Offer 26. 树的子结构
难度：  中等
编写日期：   2022 年 1 月 24 日 晚19：30
重写日期：  
*/
/*题目：  
    输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

    B是A的子结构， 即 A中有出现和B相同的结构和节点值。

    例如:
    给定的树 A:

         3
        / \
       4   5
      / \
     1   2
    给定的树 B：

       4 
      /
     1
    返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

    示例 1：

    输入：A = [1,2,3], B = [3,1]
    输出：false
    示例 2：

    输入：A = [3,4,5,1,2], B = [4,1]
    输出：true
    限制：

    0 <= 节点个数 <= 10000
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
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if B == nil {
        return false
    }
    if isEqualSub(A, B) {
        return true
    }
    flagL, flagR := false, false
    if A.Left != nil {
        flagL = isSubStructure(A.Left, B)
    }
    if A.Right != nil {
        flagR = isSubStructure(A.Right, B)
    }
    return  flagL || flagR
}

func isEqualSub(A *TreeNode, B *TreeNode) bool {
    if B == nil {
        return true
    }
    if A == nil {
        return false
    }
    if A.Val == B.Val {
        return isEqualSub(A.Left, B.Left) && isEqualSub(A.Right, B.Right)
    } else {
        return false
    }
    return false
}