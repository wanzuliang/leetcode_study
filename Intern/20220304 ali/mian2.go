// 加强编码能力

package main
 
import (
    "fmt"
    "errors"
)
 

type Queue interface {
    Enqueue(v int)
    Dequeue() (int, error)
}

type QueueNode struct {
    Val int
    Next *QueueNode
}

func (queue *QueueNode) Enqueue(data int) {

    for queue.Next != nil {
        queue = queue.Next
    }
    newNode := new(QueueNode)
    newNode.Val = data
    queue.Next = newNode
}
 
func (queue *QueueNode) Dequeue() (int, error) {
    if queue.Next == nil || queue == nil {
        return 0, errors.New("null queue")
    }
    res := queue.Next.Val
    queue.Next = queue.Next.Next
    return res, nil
}

func (queue *QueueNode) Show() {
    if queue.Next == nil {
        fmt.Println("null queue")
    }
    for queue.Next != nil {
        queue = queue.Next
        fmt.Printf(" %d ", queue.Val)
    }
    fmt.Println("")
}

func main() {
    queue := new(QueueNode)
    queue.Enqueue(7)
    queue.Enqueue(9)
    queue.Enqueue(33)
    queue.Show()

    a, err := queue.Dequeue()
    fmt.Println(a, err)
    queue.Show()


    a, err = queue.Dequeue()
    fmt.Println(a, err)
    queue.Show()


    a, err = queue.Dequeue()
    fmt.Println(a, err)
    queue.Show()


    a, err = queue.Dequeue()
    fmt.Println(a, err)
    queue.Show()
}
