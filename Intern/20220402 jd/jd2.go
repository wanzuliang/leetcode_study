package main

import (
	"fmt"
)

func main() {
	var T, N int
	for {
		t1, _ := fmt.Scan(&T)
		if t1 == 0 {
			return
		} else {
			str := make([]string, T)
			for i := 0; i < T; i++ {
				fmt.Scan(&N)
				fmt.Scan(&str[i])
				jd2(str[i])
			}
		}
	}
}
func jd2(str string) {
	res := 0
	a, b, c := 0, 0, 0
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case 'A':
			a++
		case 'B':
			b++
		case 'C':
			c++
		}
	}
	n := len(str) / 3
	if a == n && b == n && c == n {
		res = 0
		return
	}
	if a < n && b < n {
		a = n - a
		b = n - b
		res = a + b
	}
	res = a%n + b%n + c%n
	fmt.Println(res)
}

func jd1(X uint64, str string) {
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case 'U':
			if X%2 == 0 {
				X = X / 2
			} else {
				X = (X - 1) / 2
			}
		case 'L':
			X = X * 2
		case 'R':
			X = X*2 + 1
		}
	}
	fmt.Println(X)
}
