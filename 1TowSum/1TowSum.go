// see: https://leetcode.com/problems/two-sum/description/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
    result := make([]int, 2)
    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                result[0], result[1] = i, j
                return result
            }
        }

    }

    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 22

    fmt.Println(twoSum(nums, target))
}
