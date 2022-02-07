/*
题目描述：  剑指 Offer 40. 最小的k个数
难度：  简单
编写日期：   2022 年 2 月 7 日 晚 20：00
重写日期：  
*/
/*题目：  
    输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

    示例 1：

    输入：arr = [3,2,1], k = 2
    输出：[1,2] 或者 [2,1]
    示例 2：

    输入：arr = [0,1,2,1], k = 1
    输出：[0]

    限制：

    0 <= k <= arr.length <= 10000
    0 <= arr[i] <= 10000
 
*/
/*  

*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( "T")
}

// 默认代码模版
func getLeastNumbers(arr []int, k int) []int {
    for i := 0; i < k; i++ {
        for j := i+1; j < len(arr); j++ {
            if arr[j] < arr[i] {
                temp := arr[i]
                arr[i] = arr[j]
                arr[j] = temp
            }
        }
    }
    return arr[0:k]
}