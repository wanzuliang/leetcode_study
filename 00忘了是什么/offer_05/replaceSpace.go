package offer

func replaceSpace(s string) string {
    length := len(s)
    newLen := length
    // for i := 0; i < length; i++ {
    //     if s[i] == ' ' {
    //         newLen++
    //     }
    // }
    for _, v := range s {
        if v == ' ' {
            newLen += 2
        }
    }
    newStr := make([]byte, newLen)
    for i, j := length-1, newLen-1; i >= 0 ; i-- {
        if s[i] == ' ' {
            newStr[j] = '0'
            newStr[j-1] = '2'
            newStr[j-2] = '%'
            j = j-3
        } else {
            newStr[j] = s[i]
            j--
        }
    }
    return string(newStr)
}