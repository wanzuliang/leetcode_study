/*
题目描述：  剑指 Offer 06. 从尾到头打印链表
难度：  简单
编写日期：  2022 年 1 月 14 日 午15：30
重写日期：
*/
/*题目：  
    输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
    示例 1：

	输入：head = [1,3,2]
	输出：[2,3,1]
	 
	限制：
	0 <= 链表长度 <= 10000
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

type ListNode struct {
   Val int
   Next *ListNode
}

func reversePrint(head *ListNode) []int {
    if head == nil{
        return nil
    }
    list := make([]int, 0)
    return reverseInt(head, list)
}

func reverseInt(node *ListNode, list []int) []int{
    if node == nil {
        return list
    }
    list = reverseInt(node.Next, list)
    list = append(list, node.Val)
    return list
}