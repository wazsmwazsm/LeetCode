// see https://leetcode.com/problems/number-complement/description/

package main

import "fmt"

func findComplement(num int) int {
	mask := ^0
	for num & mask != 0 {
		mask <<= 1
	}
	return ^num & ^mask 
}

func main() {
	fmt.Printf("%b",findComplement(9))
}
