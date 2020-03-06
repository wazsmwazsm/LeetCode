// see: https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/

package main

import "fmt"

// 此题为有序数组
// 利用有序数组的特点，双指针一次遍历解决
// 时间 O(n) 空间 O(1)
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	min := 0
	max := len(nums) - 1

	for min < max {
		if nums[min]+nums[max] == target {
			return []int{min + 1, max + 1}
		}

		if nums[min]+nums[max] < target {
			min++
		} else {
			max--
		}

	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 22

	fmt.Println(twoSum(nums, target))
}
