/*
题目描述：  剑指 Offer 22. 链表中倒数第k个节点
难度：  简单
编写日期：   2022 年 1 月 30 日 晚 21：00
重写日期：  
*/
/*题目：  
    输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。

    例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。

    示例：

    给定一个链表: 1->2->3->4->5, 和 k = 2.

    返回链表 4->5.
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
func getKthFromEnd(head *ListNode, k int) *ListNode {
    start, end := head, head
    for i := 0; i < k; i++ {
        end = end.Next
    }
    for ; end != nil; start, end = start.Next, end.Next {
    }
    return start
}
