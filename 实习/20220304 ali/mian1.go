//给机会了，但没把握住！

//评测题目: 无

package main

import (
	"errors"
	"fmt"
)

type Queue interface {
	Enqueue(v int)
	Dequeue() (int, error)
}

type que struct {
	Val  []int
	Head int
	Last int
	Len  int
	Max  int
}

func newQ(k int) que {
	return que{
		Val:  make([]int, k),
		Head: -1,
		Last: -1,
		Len:  0,
		Max:  k,
	}
}

func (this *que) Enqueue(v int) {
	this.Len++
	if this.Max < this.Len {
		temp := make([]int, this.Max * 2)
		copy(temp, this.Val)
		this.Val = temp
		this.Max = this.Max * 2
	}
	this.Last++
	this.Last = this.Last % this.Max
	if this.Head == -1 {
		this.Head++
	}
	this.Val[this.Last] = v
	return
}

func (this *que) Dequeue() (int, error) {
	if this.Len <= 0 {
		return 0, errors.New("no data")
	} else {
		res := this.Val[this.Head]
		this.Val[this.Head] = 0
		this.Head++
		this.Len--
		this.Head = this.Head % this.Max
		return res, nil
	}
}

func (this *que) Show() {
    fmt.Printf("\nLen : %d  max : %d \n", this.Len, this.Max)
    for i := this.Head; i < this.Head+this.Len; i++ {
		fmt.Printf(" %d ", this.Val[i % this.Max])
    }
  // 	for i := 0; i < this.Max; i++ {
		// fmt.Printf(" %d ", this.Val[i])
  //   }
}

func main() {
	queue := newQ(5)
	queue.Show()

	queue.Enqueue(4)
	queue.Show()
	queue.Enqueue(3)
	queue.Show()
	queue.Enqueue(2)
	queue.Show()
	queue.Enqueue(1)
	queue.Show()

	queue.Dequeue()
	queue.Show()
	queue.Dequeue()
	queue.Show()
	queue.Dequeue()
	queue.Show()
	queue.Dequeue()
	queue.Show()

	queue.Enqueue(2)
	queue.Show()
	queue.Enqueue(1)
	queue.Show()

}
