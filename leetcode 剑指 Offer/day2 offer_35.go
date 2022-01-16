/*
题目描述：  剑指 Offer 35. 复杂链表的复制
难度：  中等
编写日期：  2022 年 1 月 14 日 晚21：00
重写日期：
*/
/*题目：  
    请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

    示例 1：

    输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
    输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
    示例 2：

    输入：head = [[1,1],[2,1]]
    输出：[[1,1],[2,1]]
    示例 3：

    输入：head = [[3,null],[3,0],[3,null]]
    输出：[[3,null],[3,0],[3,null]]
    示例 4：

    输入：head = []
    输出：[]
    解释：给定的链表为空（空指针），因此返回 null。

    提示：

    -10000 <= Node.val <= 10000
    Node.random 为空（null）或指向链表中的节点。
    节点数目不超过 1000 。
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
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    newList := copyList(head)
    randomSet(head, newList)
    return newList
}

func copyList(head *Node) *Node {
    if head == nil {
        return nil
    }
    nextNode := copyList(head.Next)
    node := &Node{Val: 1, Next: nextNode, Random: nil}
    if nextNode != nil {
        node.Val = nextNode.Val+1
    }
    return node
}

func randomSet(head *Node, newList *Node) {
    new := newList
    for node := head; node != nil; node = node.Next {
        // randomSet(head.Next, newList.Next)
        i := 0;//倒数第几个
        for l := node.Random; l != nil; l = l.Next {
            i++
        }
        
        for temp := newList; temp != nil; temp = temp.Next {
            if i == temp.Val {
                new.Random = temp
                break
            }
        }
        new = new.Next
    }
    new = newList
    for node := head; node != nil; node = node.Next {
        new.Val = node.Val
        new = new.Next
    }
}