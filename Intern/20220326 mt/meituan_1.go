package main

import (
	"fmt"
)

func main() {
	for {
		str := ""
		t1, _ := fmt.Scan(&str)
		if t1 == 0 {
			return
		} else {
			result(str)
		}
	}
}


func result(str string) {
	if len(str) < 6 {
		fmt.Println(0)
		return
	}
	a, b, c := 0, 0, 0
	for i := 0; i < len(str); i++ {
		switch str[i] {
		 case 'a':
		 	a++
		 case 'b':
		 	b++
		 case 'c':
		 	c++
		}
	}
	res := 0
	a--
	for a >= 1 && b >= 1 && c >= 3 {
		a -= 1
		b -= 1
		c -= 3
		res++
	}
	fmt.Println(res)
}

// 给你一个只包含小写字符的字符串s，你可以按任意顺序重排这个字符串中的字符，
// 请问重排过后的字符串中，最多能有多少个’acbcca’子串？

// 例如，字符串’dacbccab’中含1个’acbcca’子串，
// 按其他方式重排后最多也只能包含1个’acbcca’子串；
// 字符串’acbccaacccb’中含1个’acbcca’子串，
// 但重排成’acbccacbcca’ 后包含了2个’acbcca’子串。
// acbccaacccb
// acbccacbcca


acbccacbccacbcca ....
acbcca
	 acbcca
	 	  acbcca ....
a cbcca cbcca cbcca ....
