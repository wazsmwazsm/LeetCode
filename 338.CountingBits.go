// see https://leetcode.com/problems/counting-bits/description/
package main

import "fmt"

func countBits(num int) []int {
    count := make([]int, num + 1)
    
    for i := 0; i <= num; i++ {
        count[i] = count[i >> 1] + i & 1
    }
    
    return count
}

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
	fmt.Println(countBits(8))
}
