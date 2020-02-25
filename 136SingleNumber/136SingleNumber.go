// see https://leetcode.com/problems/single-number/description/
package main

import "fmt"

/*
func singleNumber(nums []int) int {
    m := make(map[int]int)
    
    for _, v := range nums {
        _, ok := m[v]
        if ok {
            m[v]++ 
        } else {
            m[v] = 1
        }
    }
    
    for i, v := range m {
        if v == 1 {
            return i
        }
    }
    
    return 0
}
*/

// 这个解法和巧，利用了其它都成双 (异或可抵消)，一个单独 (异或可留下)
// 当然，如果没有剩下元素必须成双这个条件，这个方法则不适应，应使用上面的方法
func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result = result ^ v 
	}
	return result
}

func main() {
	s := []int{2, 2, 1, 1, 3, 6, 3}
	fmt.Println(singleNumber(s))
}
