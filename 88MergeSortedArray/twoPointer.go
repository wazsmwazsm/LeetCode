// https://leetcode-cn.com/problems/merge-sorted-array/
package main

// 使用双指针, 利用两个数组有序、数组 1 有足够空间的特点，从后向前依次对比，比较大小后大的放到合并数组(nums1) 的最后,
// 如果数组 1 先对比完，则需将数组 2 中剩下的元素（小的元素）copy 到数组 1 中
// 如果数组 2 先对比完，因为数组 1 是合并数组，则无须操作

// 时间 O(n+m) 空间 O(1)

// 还有从前遍历的方法，但是这样需要先保存 nums1 的元素（否则会被刷新掉），增加了空间复杂度

import (
	"fmt"
)

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	fmt.Printf("num1 %v num2 %v\n", nums1, nums2)
	merge(nums1, 3, nums2, 3)
	fmt.Printf("result %v\n", nums1)

	nums1 = []int{0, 0}
	nums2 = []int{1, 2}
	fmt.Printf("num1 %v num2 %v\n", nums1, nums2)
	merge(nums1, 0, nums2, 2)
	fmt.Printf("result %v\n", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 从后往前遍历
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		// 大小对比，然后把大的往后放
		if nums1[p1] < nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}
		p--
	}

	// nums1 先遍历完后，能将未遍历的 nums2 的数据 copy 过去 (nums2 已经是小的数)
	copy(nums1[:p2+1], nums2[:p2+1])
}
