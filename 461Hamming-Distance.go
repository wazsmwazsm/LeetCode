package main

import "fmt"

func main() {

    x, y := 3, 1
    
    h := hammingDistance(x, y)
    fmt.Println(h)
}

func hammingDistance(x int, y int) int {

    count := 0
    xor := x ^ y // 异或 获得一个不同位置疑 1 的数

    for bitMask := 1; bitMask > 0; bitMask <<= 1 {
		// 使用 bitMask 移位来比较每一位是 1 还是 0
        if (bitMask & xor) != 0 {
            count++ 
        }
    }
    
    return count
}

