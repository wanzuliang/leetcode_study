// 本题为考试单行多行输入输出规范示例，无需提交，不计分。
package main

import (
	"fmt"
)

func main() {
	// a, b, x, y := 0, 0, 0, 0
	str := ""
	for {
		// n, _ := fmt.Scan(&a, &b, &x, &y)
		n, _ := fmt.Scan(&str)
		if n == 0 {
			break
		} else {
			// test1(a, b, x, y)
			_, b := test2(str)
			fmt.Println(b)
		}
	}
}

func test2(str string) (int, int) {
	// a = sum(str[i], str[i+1])
	// 			b = sum(str[i+1], str[i+2])
	// 			c = sum(str[i+2], str[i+3])
	if len(str) == 1 {
		return 0, 0
	}
	if len(str) == 2 {
		return 0, sum(str[0], str[1])
	}
	a, b := test2(str[:len(str)-1])
	aa := b
	bb := a + sum(str[len(str)-2], str[len(str)-1])
	if bb < aa {
		bb = aa
	}
	return aa, bb
}
func sum(a, b byte) int {
	if a == b {
		res := int(a-'a') + 1
		return res * 2
	} else if a-b == 1 || b-a == 1 {
		resa := int(a-'a') + 1
		resb := int(b-'a') + 1
		return resa + resb
	} else {
		return 0
	}
}
