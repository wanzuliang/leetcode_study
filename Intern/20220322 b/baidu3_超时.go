package main

import (
    "fmt"
)

/*
00101
9
9
*/
func main() {
    str := ""
    for {
        t, _ := fmt.Scan(&str)
        if t == 0 {
            break
        } else {
            temp := make(map[string]int64)
            fmt.Println(get(str, &temp) + 1)
        }
    }
}

func get(str string, m *map[string]int64) int64 {
    n := len(str)
    if n == 0 {
        return int64(0)
    }
    if n == 1 {
        return int64(1)
    }
    temp := make(map[string]string)
    for i := 0; i < n-2; i++ {
        if str[i] <= str[i+1] {
            temp[str[:i]+str[i+1:]] = str[:i]+str[i+1:]
        } else {
            temp[str[:i+1]+str[i+1+1:]] = str[:i+1]+str[i+1+1:]
        }
    }
    var res int64
    // res += 1
    for k := range temp {
        res ++
        var num int64
        if _, ok := (*m)[k]; ok {
            num = (*m)[k]
        } else {
            num = get(k, m)
        }
        res += num
    }
    return res % 1000000007
}