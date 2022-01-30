/*
题目描述：  剑指 Offer 25. 合并两个排序的链表
难度：  简单
编写日期：   2022 年 1 月 30 日 晚 21：00
重写日期：  
*/
/*题目：  
    输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

    示例1：

    输入：1->2->4, 1->3->4
    输出：1->1->2->3->4->4
    限制：

    0 <= 链表长度 <= 1000
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
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    var head, node *ListNode
    for temp1, temp2 := l1, l2; ; {
        if head == nil {
            if l1.Val <= l2.Val {
                head = l1
                temp1 = temp1.Next
            } else {
                head = l2
                temp2 = temp2.Next
            }
            node = head
        }
        if temp1 == nil {
            node.Next = temp2
            break
        }
        if temp2 == nil {
            node.Next = temp1
            break
        }
        if temp1.Val <= temp2.Val {
            temp := temp1.Next 
            node.Next = temp1
            node = node.Next
            temp1 = temp
        } else {
            temp := temp2.Next 
            node.Next = temp2
            node = node.Next
            temp2= temp
        }
    }
    return head
}