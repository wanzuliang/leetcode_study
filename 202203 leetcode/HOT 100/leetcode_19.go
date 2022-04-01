/*
题目描述：  19. 删除链表的倒数第 N 个结点
难度：  中等
编写日期：   2022 年 3 月 23 日 21：30
重写日期：   
*/
/*题目：  
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

 

示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]
 

提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
 

进阶：你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    temp1, temp2 := head, head			//双指针搁 n 个位置
    if head.Next == nil && n == 1 {		//当只有一个节点时
        return nil
    }
    for n > 0 {							//尾指针先走 n 个位置
        n--
        temp2 = temp2.Next
    }
    if temp2 == nil {					//如果尾指针到尽头了，正好倒数 n 个值是头节点
        head = temp1.Next	//	去掉第一个节点
        return head
    }
    for temp2.Next != nil {				//正常让两个指针同步走，直到尾指针到尽头
        temp1 = temp1.Next
        temp2 = temp2.Next
    }
    temp1.Next= temp1.Next.Next	//删除节点
    return head
}