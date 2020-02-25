// see https://leetcode-cn.com/problems/flip-game/
package main

import (
	"fmt"
)

var tests = []string{
	"+++--++-",
	"++++-++-+--+-++-",
	"++++",
	"--",
	"-+-++",
}

func main() {
	for _, test := range tests {
		result := generatePossibleNextMoves(test)
		fmt.Println(result)
	}
}

func generatePossibleNextMoves(s string) []string {
	result := []string{}

	for i := 0; i < len(s)-1; i++ {
		tmp := []byte(s)
		if s[i] == '+' && s[i+1] == '+' {
			tmp[i] = '-'
			tmp[i+1] = '-'
			result = append(result, string(tmp))
		}
	}

	return result
}
