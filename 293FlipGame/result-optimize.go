// see https://leetcode-cn.com/problems/flip-game/
package main

import (
	"fmt"
	"strings"
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
	var buf strings.Builder
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '+' && s[i+1] == '+' {
			buf.WriteString(s[0:i])
			buf.WriteString("--")
			buf.WriteString(s[i+2:])
			result = append(result, buf.String())
			buf.Reset()
		}
	}

	return result
}
