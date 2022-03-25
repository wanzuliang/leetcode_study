/*
题目描述：  21. 合并两个有序链表
难度：  简单
编写日期：   2022 年 3 月 25 日 21：30
重写日期：   
*/
/*题目：  
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

 

示例 1：


输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]
 

提示：

两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println("T")
}

// 默认代码模版
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 /**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    res := new(ListNode)
    t1, t2 := list1, list2
    t := res
    for {
        if t1 == nil && t2 == nil{
            break
        } else {
            if t1 == nil {
                node := new(ListNode)
                node.Val = t2.Val
                node.Next = nil
                t.Next = node
                t = node
                t2 = t2.Next
                continue
            }
            if t2 == nil {
                node := new(ListNode)
                node.Val = t1.Val
                node.Next = nil
                t.Next = node
                t = node
                t1 = t1.Next
                continue
            }
            if t1.Val < t2.Val {
                node := new(ListNode)
                node.Val = t1.Val
                node.Next = nil
                t.Next = node
                t = node
                t1 = t1.Next
            } else {
                node := new(ListNode)
                node.Val = t2.Val
                node.Next = nil
                t.Next = node
                t = node
                t2 = t2.Next
            }
        }
    }
    return res.Next
}