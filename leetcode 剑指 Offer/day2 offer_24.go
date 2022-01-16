/*
题目描述：  剑指 Offer 24. 反转链表
难度：  简单
编写日期：  2022 年 1 月 14 日 晚19：00
重写日期：
*/
/*题目：  
    定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

    示例:
    输入: 1->2->3->4->5->NULL
    输出: 5->4->3->2->1->NULL

    限制：
    0 <= 节点个数 <= 5000
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
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    B := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return B
}