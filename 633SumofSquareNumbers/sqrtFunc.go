package main

// 枚举 a, 直接开发库函数找到 b, 看看是不是整数 (对比 ceil 和 floor)
// 时间 O(sqrt(c)), 空间 O(1)

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

	for a := 0; a*a <= c; a++ {
		b := math.Sqrt(float64(c - a*a))
		if math.Ceil(b) == math.Floor(b) {
			return true
		}
	}

	return false
}
