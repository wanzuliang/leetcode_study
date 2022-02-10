/*
题目描述：  剑指 Offer 41. 数据流中的中位数
难度：  困难
编写日期：   2022 年 2 月 7 日 晚 20：00
重写日期：   2022 年 2 月 10 日 晚 20：00 
*/
/*题目：  
    如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

    例如，

    [2,3,4] 的中位数是 3

    [2,3] 的中位数是 (2 + 3) / 2 = 2.5

    设计一个支持以下两种操作的数据结构：

    void addNum(int num) - 从数据流中添加一个整数到数据结构中。
    double findMedian() - 返回目前所有元素的中位数。
    示例 1：

    输入：
    ["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
    [[],[1],[2],[],[3],[]]
    输出：[null,null,null,1.50000,null,2.00000]
    示例 2：

    输入：
    ["MedianFinder","addNum","findMedian","addNum","findMedian"]
    [[],[2],[],[3],[]]
    输出：[null,null,2.00000,null,2.50000]
     

    限制：

    最多会对 addNum、findMedian 进行 50000 次调用。
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
type MedianFinder struct {
    Val int
    Next *MedianFinder
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
    var head MedianFinder
    return head
}


func (this *MedianFinder) AddNum(num int)  {
    temp := this
    node := new(MedianFinder)
    node.Val = num
    node.Next = nil
    for temp.Next != nil  {
        if num < temp.Next.Val {
            node.Next = temp.Next
            temp.Next = node
            this.Val = this.Val + 1
            break
        }
        temp = temp.Next
    }
    if temp.Next == nil {
        node.Next = temp.Next
        temp.Next = node
        this.Val = this.Val + 1
    }
}


func (this *MedianFinder) FindMedian() float64 {
    temp := this
    step := 0
    for float64(step) < float64(this.Val) / 2 {
        step ++
        temp = temp.Next
    }
    // return float64(step)
    if this.Val % 2 == 0 {
        return (float64(temp.Val) + float64(temp.Next.Val)) / 2
    } else {
        return float64(temp.Val)
    }
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */