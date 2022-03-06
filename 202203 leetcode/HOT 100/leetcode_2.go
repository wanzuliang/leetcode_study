/*
题目描述：  2. 两数相加
难度：  中等
编写日期：   2022 年 3 月 6 日 22：04
重写日期：   
*/
/*题目：  

	给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

    请你将两个数相加，并以相同形式返回一个表示和的链表。

    你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

     

    示例 1：


    输入：l1 = [2,4,3], l2 = [5,6,4]
    输出：[7,0,8]
    解释：342 + 465 = 807.
    示例 2：

    输入：l1 = [0], l2 = [0]
    输出：[0]
    示例 3：

    输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
    输出：[8,9,9,9,0,0,0,1]
     

    提示：

    每个链表中的节点数在范围 [1, 100] 内
    0 <= Node.val <= 9
    题目数据保证列表表示的数字不含前导零

    来源：力扣（LeetCode）
    链接：https://leetcode-cn.com/problems/add-two-numbers
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
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        L = ListNode()
        l0 = L
        while l1 != None or l2 != None:
            temp = l0.val
            if l1 != None:
                a = l1.val
                l1 = l1.next
            else:
                a = 0
            if l2 != None:
                b = l2.val
                l2 = l2.next
            else:
                b = 0
            temp = a+b+temp
            l0.val = temp%10
            temp = temp//10
            if l1 != None or l2 !=None or temp !=0:
                l0.next = ListNode(val = temp)
                l0 = l0.next
            
        return L
        
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var res, node *ListNode
    add := 0
    temp1 := l1
    temp2 := l2
    for temp1 != nil || temp2 != nil {
        a, b := 0, 0
        var temp ListNode
        if temp1 != nil {
            a = temp1.Val 
            temp1 = temp1.Next
        }
        if temp2 != nil {
            b = temp2.Val 
            temp2 = temp2.Next
        }
        val := a + b + add 
        temp.Val = val % 10
        add = val / 10
        if res == nil {
            res = &temp
            node = &temp
        } else {
            node.Next = &temp
            node = node.Next
        }
    }
    if add == 1 {
        var temp ListNode
        temp.Val = 1
        node.Next = &temp
        node = node.Next
    }
    return res
}