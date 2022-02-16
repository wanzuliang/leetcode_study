/*
题目描述：  剑指 Offer 29. 顺时针打印矩阵
难度：  简单
编写日期：   2022 年 2 月 16 日 晚 12：00
重写日期：  
*/
/*题目：  

	输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

    示例 1：

    输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
    输出：[1,2,3,6,9,8,7,4,5]
    示例 2：

    输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
    输出：[1,2,3,4,8,12,11,10,9,5,6,7]
     

    限制：

    0 <= matrix.length <= 100
    0 <= matrix[i].length <= 100

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
func spiralOrder(matrix [][]int) []int {

    high := len(matrix)
    if high == 0 {
        return nil
    }
    width := len(matrix[0])
    if width == 0 {
        return nil
    }
    temp := make([]int,0)
    
    if high == 1 {
        for i := 0; i < width; i++ {
            temp = append(temp, matrix[0][i])
        }
        return temp
    }
    if width == 1 {
        for i := 0; i < high; i++ {
            temp = append(temp, matrix[i][0])
        }
        return temp
    }
    if high > 1 && width > 1 {
        for i := 0; i < width - 1; i++ {
            temp = append(temp, matrix[0][i])
        }
        for i := 0; i < high - 1; i++ {
            temp = append(temp, matrix[i][width-1])
        }
        for i := width - 1; i > 0; i-- {
            temp = append(temp, matrix[high-1][i])
        }
        for i := high - 1; i > 0; i-- {
            temp = append(temp, matrix[i][0])
        }
        // temp2 := spiralOrder(matrix[1:high-1][1:width-1])
        // 这样切片是连续取二维数组的行
        new_matrix := make([][]int,0)
        for i := 1; i < high-1; i++ {
            new_matrix = append(new_matrix, matrix[i][1:width-1])
        }
        temp2 := spiralOrder(new_matrix)
        for _,v := range temp2 {
            temp = append(temp, v)
        }
        return temp
    }
    return nil
}