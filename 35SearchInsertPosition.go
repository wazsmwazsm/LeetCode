// see https://leetcode.com/problems/search-insert-position/description/

package main

import "fmt"

func searchInsert(nums []int, target int) int {
    
    result := 0
    
    for i, v := range nums {

        if v == target {
            return i
        }
        
        if target > v {
            result = i + 1
        }
    }
    
    return result
    
}

func main() {
	s := []int{1, 3, 4, 7, 8}

	fmt.Println(searchInsert(s, 7))
	fmt.Println(searchInsert(s, 2))
}