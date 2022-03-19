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

        fmt.Printf("%d    %d\n", 39, sumMy(39))
    // for n := 3; n < 40; n++ {
    //     // fmt.Printf("%d    %d\n", n, sumMy(n))
    //     fmt.Printf("%d  %d  %d  %d\n", n, other(n), sumMy(n))
    // }
}

func sumMy(n int) int {
    res := 0
    all := float64(n-2) * 180.0
    eg := all / float64(n)
    now := eg
    fmt.Printf("%f  %f  %f \n", all, eg, now)
    same := 0
    for now > 0 {
        if now < 90{
            if now == 60 {
                same -= 2 * n / 3
            }
            res++
        }
        now += (eg - 180.0)
    }
    return res * n + same
}

func other(n int) (res int) {   //  https://www.nowcoder.com/discuss/854699
    res = 0
    if n % 2 == 0 {
      res = n * ((n-2) / 4);
    } else {
      res = n * (((n-1) / 2+1) / 2)
    }
    if n % 3 == 0 {
      res -= (n / 3) * 2
    }
    return
}
