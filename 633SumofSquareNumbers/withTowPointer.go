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
