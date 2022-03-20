package main
import (
    "fmt"
    "os"
    "strings"
    "bufio"
    "strconv"
)
func main() {
    input := bufio.NewReader(os.Stdin)
    var sum int
    for {
        s, _, _ := input.ReadLine()
        if len(s) == 0 {
            break
        }
        arr := string(s)
        num := strings.Fields(arr)
        sum = 0
        for _, v := range num {
            t, _ := strconv.Atoi(v)
            sum += t
        }
        fmt.Println(sum)
    }
}