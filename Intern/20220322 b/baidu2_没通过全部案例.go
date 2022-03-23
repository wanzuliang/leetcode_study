package main

import (
    "fmt"
)
func main() {
    str := ""
    for {
        t, _ := fmt.Scan(&str)
        if t == 0 {
            break
        } else {
           show(str)
        }
    }
}

func show(str string) {
    n := len(str)
    a, b, c, d := 0, 0, 0, 0
    for i := 0; i < n; i++ {
        if i == n-1 {
            if str[i] == str[n-1] {
                a, b = 0, n-1
            } else {
                a, b = 0, n-2
            }
            break
        }
        if str[i] == str[n-1] {
            a, b = i, n-1
            break
        }
    }
    for i := n-1; i >= 0; i-- {
        if i == 0 {
            if str[i] == str[0] {
                a, b = 0, n-1
            } else {
                a, b = 1, n-1
            }
            break
        }
        if str[i] == str[0] {
            c, d = i, 0
            break
        }
    }
    if b-a < c-d {
        fmt.Printf("%d %d %d %d\n", a+1, b, a+2, b+1)
    } else {
        fmt.Printf("%d %d %d %d\n", d+1, c, d+2, c+1)
    }
}