package main

import "fmt"

// 使用 strings 包的话 strings.ToLower(str) 即可
func toLowerCase(str string) string {
    new_str := make([]rune, len(str))
    for i, v := range str {
        if v >= 65 && v <= 90 {
            new_str[i] = v + 32
        } else {
            new_str[i] = v
        }
    }
    
    return string(new_str)
}

func main() {
    fmt.Println(toLowerCase("HelLo"))
}
