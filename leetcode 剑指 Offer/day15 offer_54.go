/*
题目描述：  剑指 Offer 54. 二叉搜索树的第k大节点
难度：  中等
编写日期：   2022 年 2 月 3 日 晚 21：00
重写日期：  
*/
/*题目：  
    给定一棵二叉搜索树，请找出其中第 k 大的节点的值。

    示例 1:
    输入: root = [3,1,4,null,2], k = 1
       3
      / \
     1   4
      \
       2
    输出: 4

    示例 2:
    输入: root = [5,3,6,2,4,null,null,1], k = 3
           5
          / \
         3   6
        / \
       2   4
      /
     1
    输出: 4

    限制：
    1 ≤ k ≤ 二叉搜索树元素个数
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
func kthLargest(root *TreeNode, k int) int {
    result := 0
    var recur func(root*TreeNode)
    recur = func(root*TreeNode) {
        if root == nil || k <= 0 {
            return
        }
        recur(root.Right)
        k--
        if k == 0 {
            result = root.Val
            return
        }
        recur(root.Left)
    }
    recur(root)
    return result
}