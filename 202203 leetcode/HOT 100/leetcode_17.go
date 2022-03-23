/*
题目描述：  17. 电话号码的字母组合
难度：  中等
编写日期：   2022 年 3 月 21 日 22：30
重写日期：   
*/
/*题目：  
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

	
	2	{'a', 'b', 'c'}, 
    3   {'d', 'e', 'f'}, 
    4   {'g', 'h', 'i'},
	5	{'j', 'k', 'l'}, 
    6    {'m', 'n', 'o'}, 
    7    {'p', 'q', 'r', 's'},
	8	{'t', 'u', 'v'}, 
    9    {'w', 'x', 'y', 'z'}

 

示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]
 

提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( letterCombinations("2") )
    fmt.Println( letterCombinations("29") )
    fmt.Println( letterCombinations("239") )
}

// 默认代码模版
func letterCombinations(digits string) []string {
	if len(digits) == 0 || digits[0] == '1' {
		return []string{}
	}
	str := [][]byte{
		{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'},
		{'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'},
		{'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}

	res := make([]string, 0)
	if len(digits) == 1 {
		for _, ch := range str[digits[0]-'2'] {
			res = append(res, string(ch))
		}
		return res
	}
	if len(digits) > 1 {
		befor := letterCombinations(digits[:len(digits)-1])
		last := str[digits[len(digits)-1]-'2']
		for _, ss := range befor {
			for _, ch := range last {
				res = append(res, ss+string(ch))
			}
		}
	}
	return res
}