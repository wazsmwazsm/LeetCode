// https://leetcode-cn.com/problems/sum-of-square-numbers/description/

// 双指针解法 (逼近、缩减范围)
// 时间 O(sqrt(c)) (从 0 ~ sqrt(c) 遍历) 空间 O(1)

package main

import (
	"fmt"
	"math"
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

	if c < 0 {
		return false
	}

	a := 0
	b := int(math.Sqrt(float64(c)))

	for a <= b {
		pow := a*a + b*b
		if pow == c {
			return true
		}

		if pow < c {
			a++
		} else {
			b--
		}
	}

	return false
}
