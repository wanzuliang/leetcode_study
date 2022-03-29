// 本题为考试单行多行输入输出规范示例，无需提交，不计分。
package main

import (
	"fmt"
	"math"
)

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func main() {
	a := 0
	for {
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		} else {
			test3(a)
		}
	}
}

func test3(n int) {

	// if n%2 == 1 {
	hight := 1
	for math.Pow(2, float64(hight)) <= float64(n) {
		hight++
	}
	node := int(math.Pow(2, float64(hight-1))) - 1
	sonsum := n - node
	son2sum := int(math.Pow(2, float64(hight-2))) - (sonsum+1)/2
	// son1sum := int(math.Pow(2, float64(hight-2))) - son2sum
	// fmt.Println(hight, node, sonsum, son1sum, son2sum)
	max := 0
	for i := 1; i <= node-son2sum; i++ {
		fmt.Printf("%d ", i*2)
		max = i * 2
	}
	for i := 1; i <= n; i += 2 {
		fmt.Printf("%d ", i)
	}
	for i := max + 2; i < n; i++ {
		fmt.Printf("%d ", i)
	}
	// }

}
