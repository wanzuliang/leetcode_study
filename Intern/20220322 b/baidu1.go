package main

import (
    "fmt"
)

/*扩大图像
3 3 
1 0 1 
0 0 0 
1 0 1
*/
func main() {
    n := 0
    k := 0
    for {
        t, _ := fmt.Scan(&n, &k)
        if t == 0 {
            break
        } else {
            res := make([][]int, 0)
            for i := 0; i < n; i++ {
                line := make([]int, n)
                for j := 0; j < n; j++ {
                    fmt.Scan(&line[j])
                }
                res = append(res, line)
            }
            show(res, k)
        }
    }
}

func show(pic [][]int, k int) {
    n := len(pic[0])
    for i := 0; i < n*k; i++ {
        for j := 0; j < n*k; j++ {
            fmt.Printf("%d ", pic[i/k][j/k])
        }
        fmt.Println()
    }
}