package main
import (
    "fmt"
)
// func main() {
//     // n 边行 内角和 (n-2)*180
//     // n 边行 每个角 (n-2)*180 / n
//     // - jiao/2

//     n := 0
//     sum := 0
//     for {
//         t1, _ := fmt.Scan(&n)
//         if t1 == 0 {
//             break
//         } else {
//             sum = sumMy(n)
//         }
//         fmt.Printf("%d\n", sum)
//     }
// }
func main() {
    for n := 3; n < 20; n++ {
        fmt.Printf("%d    %d\n", n, sumMy(n))
    }
}
// 3    1
// 4    0
// 5    5
// 6    2
// 7    14
// 8    8
// 9    3
// 10    20
// 11    33
// 12    4
// 13    39
// 14    42
// 15    5
// 16    64
// 17    68
// 18    6
// 19    95

func sumMy(n int) int {
    if n % 3 == 0 {
        return n / 3
    }
    res := 0
    all := (n-2) * 180.0
    eg := all / n * 1.0
    now := eg * 1.0
    for now > 0 {
        if now >= 90 {
            now -= (180.0 - eg)
            continue
        } else {
            res++
            now -= (180.0 - eg)
        }
    }
    return res * n 
}