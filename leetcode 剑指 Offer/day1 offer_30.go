/*
题目描述：  剑指 Offer 30. 包含min函数的栈
难度：  简单
编写日期：  2022 年 1 月 13 日 晚21：45
重写日期：
*/
/*题目：  
    定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
    示例:
    MinStack minStack = new MinStack();
    minStack.push(-2);
    minStack.push(0);
    minStack.push(-3);
    minStack.min();   --> 返回 -3.
    minStack.pop();
    minStack.top();      --> 返回 0.
    minStack.min();   --> 返回 -2.
    提示：
    各函数的调用总次数不超过 20000 次
*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println("T")
    obj := Constructor();
    obj.Push(-1);
    obj.Push(0);
    obj.Push(-2);
    obj.Pop();
    param_3 = obj.Top();
    param_4 = obj.Min();
    fmt.Println("Top: ", param_3)
    fmt.Println("Min: ", param_4)
}

// 填写默认代码模版
type MinStack struct {
    min []int
    stack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{min: make([]int, 0),stack: make([]int, 0)}
}


func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    if len(this.min) == 0 || x <= this.Min(){
        this.min = append(this.min, x)
    // } else if  {
    //     this.min = append(this.min, x)
    }
}


func (this *MinStack) Pop()  {
    if this.Top() == this.Min() {
        this.min = this.min[:len(this.min) - 1]
    }
    this.stack = this.stack[:len(this.stack) - 1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack) - 1]
}


func (this *MinStack) Min() int {
    return this.min[len(this.min) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */