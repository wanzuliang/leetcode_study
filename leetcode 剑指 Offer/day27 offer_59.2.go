/*
题目描述：  剑指 Offer 59 - II. 队列的最大值
难度：  中等
编写日期：   2022 年 2 月 18 日 晚 20：00
重写日期：  
*/
/*题目：  

	请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

	若队列为空，pop_front 和 max_value 需要返回 -1

	示例 1：

	输入: 
	["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
	[[],[1],[2],[],[],[]]
	输出: [null,null,null,2,1,2]
	示例 2：

	输入: 
	["MaxQueue","pop_front","max_value"]
	[[],[],[]]
	输出: [null,-1,-1]
	 

	限制：

	1 <= push_back,pop_front,max_value的总操作数 <= 10000
	1 <= value <= 10^5

*/
/*  
		重写
*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( "T")
}

// 默认代码模版
// type MaxQueue struct {
//     Queue []int
//     Len int
//     Max int
//     Next *MaxQueue
//     Last *MaxQueue
// }


// func Constructor() MaxQueue {
//     var queue MaxQueue
//     return queue
// }


// func (this *MaxQueue) Max_value() int {
//     if this.Len == 0 {
//         return -1
//     } else {
//         return this.Max
//     }
// }


// func (this *MaxQueue) Push_back(value int) {
//     if this.Last == nil {
//         if this.Len == 0 {
//             this.Queue = append(this.Queue, value)
//             this.Max = value
//             this.Len = this.Len + 1
//         } else if value >= this.Max {
//             this.Queue = append(this.Queue, value)
//             this.Max = value
//             this.Len = this.Len + 1
//         } else if value < this.Max {
//             temp := new(MaxQueue)
//             temp.Queue = append(temp.Queue, value)
//             temp.Max = 1
//             this.Len = this.Len + 1
//             this.Next = temp
//             this.Last = temp
//         }
//     } else {
//         lastQueue := this.Last
//         if value >= this.Max {
//             lastQueue.Queue = append(lastQueue.Queue, value)
//             lastQueue.Max = value
//             this.Len = this.Len + 1
//         } else if value < this.Max {
//             temp := new(MaxQueue)
//             temp.Queue = append(temp.Queue, value)
//             temp.Max = value
//             this.Len = this.Len + 1
//             this.Next = temp
//             this.Last = temp
//         }
//     }
// }


// func (this *MaxQueue) Pop_front() int {
//     if this.Len == 0 {
//         return -1
//     } else {
//         if len(this.Queue) == 0 {
//             this.Next.Last = this.Last
//             this.Next.Len = this.Len
//             this = this.Next
//         }
//         this.Len--
//         res := this.Queue[0]
//         this.Queue = this.Queue[1:]
//         return res
//     }
// }

type MaxQueue struct {
    q []int
    p []int
}

func Constructor() MaxQueue {
    return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
    if len(this.q) == 0 {
        return -1
    }
    return this.p[0]
}

func (this *MaxQueue) Push_back(value int)  {
    this.q = append(this.q, value)
    for len(this.p) > 0 && value > this.p[len(this.p) - 1] {
        this.p = this.p[: len(this.p) - 1]
    }
    this.p = append(this.p, value)
}

func (this *MaxQueue) Pop_front() int {
    if len(this.q) == 0 {
        return -1
    }
    if this.q[0] == this.p[0] {
        this.p = this.p[1: ]
    }
    value := this.q[0]
    this.q = this.q[1: ]
    return value
}
/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */