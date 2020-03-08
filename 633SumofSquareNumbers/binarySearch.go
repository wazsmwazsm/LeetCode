package main

// 二分查找法, 枚举 a, 二分查找 b
// 时间 O(log(c)sqrt(c)) (二分是 log(c) 枚举是 sqrt(c)), 空间 O(log(c)) 非递归 O(1)

import (
	"fmt"
)

func main() {
	inputs := map[int]bool{
		3: false,
		5: true,
		6: false,
		7: false,
		8: true,
		9: true,
	}

	for input, res := range inputs {
		if judgeSquareSum(input) != res {
			panic("error")
		}
	}

	fmt.Println("Done")
}

func judgeSquareSum(c int) bool {

	for a := 0; a*a <= c; a++ {
		b := c - a*a
		// if binarySearch(0, b, b) {
		if binarySearch2(0, b, b) {
			return true
		}
	}

	return false
}

func binarySearch(s, e, n int) bool {
	if s > e {
		return false
	}

	mid := s + int((e-s)/2)

	if mid*mid == n {
		return true
	} else if n > mid {
		return binarySearch(mid+1, e, n)
	}

	return binarySearch(s, mid-1, n)
}

func binarySearch2(s, e, n int) bool {

	for s <= e {
		mid := s + int((e-s)/2)
		if mid*mid == n {
			return true
		} else if n > mid*mid {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}

	return false
}
