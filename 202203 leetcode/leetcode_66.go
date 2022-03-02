/*
题目描述：  66. 加一
难度：  简单
编写日期：   2022 年 3 月 2 日 22：30
重写日期：   
*/
/*题目：  

	给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

	最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

	你可以假设除了整数 0 之外，这个整数不会以零开头。

	 

	示例 1：

	输入：digits = [1,2,3]
	输出：[1,2,4]
	解释：输入数组表示数字 123。
	示例 2：

	输入：digits = [4,3,2,1]
	输出：[4,3,2,2]
	解释：输入数组表示数字 4321。
	示例 3：

	输入：digits = [0]
	输出：[1]
	 

	提示：

	1 <= digits.length <= 100
	0 <= digits[i] <= 9

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/plus-one
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( "T" )
}

// 默认代码模版
func plusOne(digits []int) []int {
    add := 1
    for i := len(digits) - 1; i >= 0 ; i-- {
        if digits[i] == 9 && add == 1 {
            if i == 0 {
                temp := make([]int, 1 + len(digits))
                temp[0] = 1
                return temp
            }
            digits[i] = 0
            add = 1
        } else {
            digits[i] += add
            add = 0
        }
    }
    return digits
}