/*
题目描述：  15. 三数之和
难度：  中等
编写日期：   2022 年 3 月 15 日 22：30
重写日期：   
*/
/*题目：  
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

 

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]
 

提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
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
func longestPalindrome(s string) string {
    max := 1
    l, r := 0, 0
    for n := 0; n < len(s); n++ {
        for i := 1; 0 <= n-i && n+i < len(s); i++ { //奇数回文
            if s[n-i] == s[n+i] {
                if i*2+1 > max {
                    max = i*2 + 1
                    l = n - i
                    r = n + i
                }
            } else {
                break
            }
        }
        for i := 1; 0 <= n-i+1 && n+i < len(s); i++ { //偶数回文
            if s[n-i+1] == s[n+i] {
                if 2*i > max {
                    max = 2 * i
                    l = n - i + 1
                    r = n + i
                }
            } else {
                break
            }
        }
    }
    return s[l:r+1]
}