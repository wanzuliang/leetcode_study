/*
题目描述：  剑指 Offer 09. 用两个栈实现队列
难度：  简单
编写日期：  2022 年 1 月 13 日 晚21：00
重写日期：
*/
/*题目：  
    用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
    分别完成在队列尾部插入整数和在队列头部删除整数的功能。
    (若队列中没有元素， deleteHead 操作返回 -1 )
    示例 1：
        输入：
        ["CQueue","appendTail","deleteHead","deleteHead"]
        [[],[3],[],[]]
        输出：[null,null,3,-1]
    示例 2：
        输入：
        ["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
        [[],[],[5],[2],[],[]]
        输出：[null,-1,null,null,5,2]
        提示：

    1 <= values <= 10000
    最多会对 appendTail 、 deleteHead 进行 10000 次调用
*/
package main

import (
    "fmt"
)

func main() {
    obj := Constructor()

    param_2 := obj.DeleteHead()
    fmt.Println(param_2)

    obj.AppendTail(3)

    param_2 = obj.DeleteHead()
    fmt.Println(param_2)

    param_2 = obj.DeleteHead()
    fmt.Println(param_2)

}

// 填写默认代码模版
type CQueue struct {
    stackIn []int
    stackOut []int  
}

func Constructor() CQueue {
    return CQueue{stackIn: make([]int, 0), stackOut: make([]int, 0)}
}

func (this *CQueue) AppendTail(value int)  {
    this.stackIn = append(this.stackIn, value)
}

func (this *CQueue) DeleteHead() int {
    if len(this.stackOut) == 0 && len(this.stackIn) == 0{
        return -1
    }

    if len(this.stackOut) == 0 {
        stackInLen := len(this.stackIn)
        for i := stackInLen-1; i > 0; i-- {
            this.stackOut = append(this.stackOut, this.stackIn[i])
        }
        out := this.stackIn[0]
        this.stackIn = this.stackIn[:0]
        return out
    }

    stackOutLen := len(this.stackOut)
    out := this.stackOut[stackOutLen-1]
    this.stackOut = this.stackOut[:stackOutLen-1]
    return out
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
