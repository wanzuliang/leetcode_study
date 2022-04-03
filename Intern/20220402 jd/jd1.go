package main

import (
	"fmt"
	"math/big"
)

func main() {
	var N, X big.Int
	for {
		t1, _ := fmt.Scan(&N, &X)
		if t1 == 0 {
			return
		} else {
			str := ""
			fmt.Scan(&str)
			jd1(X, str)
		}
	}
}

func jd1(X big.Int, str string) {
	var a *big.Int
	a = big.NewInt(2)
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case 'U':
			if X%int(a) == 0 {
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
