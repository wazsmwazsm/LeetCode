// see https://leetcode.com/problems/find-the-duplicate-number/description/
package main

import "fmt"

func findDuplicate(nums []int) int {
    m := make(map[int]int) 
    
    for _, v := range nums {
        _, ok := m[v] 
        if ok {
            return v
        } else {
            m[v] = 1 
        }
    }
    
    return 0
}

func main() {
	s1 := []int {1,3,4,2,2}
	s2 := []int {3,1,3,4,2}
	fmt.Println(findDuplicate(s1))
	fmt.Println(findDuplicate(s2))
}
