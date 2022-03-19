package main
import (
    "fmt"
)
func main() {
    n := 0
    a := 0
    b := 0
    sum := 0
    for {
        t1, _ := fmt.Scan(&n)
        if t1 == 0 {
            break
        } else {
            for n > 0 {
                t2, _ := fmt.Scan(&a,&b)
                if t2 == 0 {
                    break
                } else {
                    sum += sumMy(a, b)
                }
            }
            
        }
        fmt.Printf("%d\n", sum)
    }
}

func sumMy(a, b int) int {
    if b == 0 {
        return 1
    }
    res := sumMy(a, b - 1) +  b * (a - 2) + 1
    return res
}