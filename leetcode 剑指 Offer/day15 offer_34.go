/*
题目描述：  剑指 Offer 34. 二叉树中和为某一值的路径
难度：  中等
编写日期：   2022 年 2 月 2 日 晚 20：00
重写日期：  
*/
/*题目：  
    给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

    叶子节点 是指没有子节点的节点。

    示例 1：

    输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
    输出：[[5,4,11,2],[5,8,4,5]]
    示例 2：

    输入：root = [1,2,3], targetSum = 5
    输出：[]
    示例 3：

    输入：root = [1,2], targetSum = 0
    输出：[]

    提示：

    树中节点总数在范围 [0, 5000] 内
    -1000 <= Node.val <= 1000
    -1000 <= targetSum <= 1000
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
func pathSum(root *TreeNode, target int) [][]int {          //通过不了，暂时不知道为什么错了
    if root == nil {
        return nil
    }
    result := [][]int{}
    road := []int{}
    dfs(root, target, road, &result)
    return result
}

func dfs(root *TreeNode, target int, road []int, result *[][]int) {
    road = append(road, root.Val)
    if root.Right == nil && root.Left == nil && root.Val == target {
        *result = append(*result, road)
        return
    } else {
        if root.Left != nil {
            dfs(root.Left , target - root.Val, road, result)
        }
        if root.Right != nil {
            dfs(root.Right, target - root.Val, road, result)
        }
    }
    return
}

// 作者：user8744q
// 链接：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/solution/go-jian-zhi-offer-34-er-cha-shu-zhong-he-de3y/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
    if root == nil {
        return [][]int{}
    }
    ans := [][]int{}
    var dfs func(root *TreeNode, cur int, tmp []int)
    dfs = func(root *TreeNode, cur int, tmp []int) {
        cur += root.Val
        tmp = append(tmp, root.Val)
        if root.Left == nil && root.Right == nil{
            if target == cur {
                //相当于建立临时slice再copy
                // t := make([]int, len(tmp))
                // copy(t, tmp)
                // ans = append(ans, t)
                //为什么这么做？
                //始终谨记go的参数传递为值传递，就是copy的副本
                //错误做法
                //仅仅 ans = append(ans, tmp)
                //实际上我们只需要保存每一层的slice即可
                ans = append(ans, append([]int(nil), tmp...))
                return
            }
        }
        if root.Left != nil {dfs(root.Left, cur, tmp)}
        if root.Right != nil {dfs(root.Right, cur, tmp)}
    }
    dfs(root, 0, []int{})

    return ans
}

