/*
题目描述：  剑指 Offer 63. 股票的最大利润
难度：  中等
编写日期：   2022 年 1 月 25 日 晚 20：00
重写日期：  
*/
/*题目：  
    假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

    示例 1:
    输入: [7,1,5,3,6,4]
    输出: 5
    解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
         注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

    示例 2:
    输入: [7,6,4,3,1]
    输出: 0
    解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

    限制：
    0 <= 数组长度 <= 10^5
*/
/*  
    
*/
package main

import (
    "fmt"
)

// func main() {
//     fmt.Println( "T")
// }

// 默认代码模版
func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    min, max := prices[0], prices[0]
    get := max - min
    for i := 0; i < len(prices); i++ {
        if prices[i] > max {
            if prices[i] - min > get {
                get = prices[i] - min
            }
            max = prices[i] 
        } else if prices[i] < min {
            min = prices[i]
            max = prices[i]
        }
    }
    return get
}