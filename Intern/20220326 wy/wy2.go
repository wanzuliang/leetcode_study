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
			test2(str)
		}
	}
}

func test2(str string) {
	// n, m := 0, 0
	res := 0
	for i := 0; i < len(str); {
		a, b, c := 0, 0, 0
		if len(str)-i <= 4 {
			if len(str)-i == 2 {
				res += sum(str[i], str[i+1])
			}
			if len(str)-i == 3 {
				a = sum(str[i], str[i+1])
				b = sum(str[i+1], str[i+2])
				if a > b {
					res += a
				} else {
					res += b
				}
			}
			if len(str)-i == 4 {
				a = sum(str[i], str[i+1])
				b = sum(str[i+1], str[i+2])
				c = sum(str[i+2], str[i+3])
				if b >= a+c {
					res += b
				} else {
					res += (a + c)
				}
			}
			break
		} else {
			a = sum(str[i], str[i+1])
			b = sum(str[i+1], str[i+2])
			c = sum(str[i+2], str[i+3])
			if b >= a+c {
				res += b
				i += 3
			} else if a > 0 {
				res += a
				i += 2
			} else {
				i++
			}
		}

	}
	fmt.Println(res)
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
