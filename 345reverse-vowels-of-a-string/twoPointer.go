package main

// 双指针. 找到元音字母则替换

// 时间 O(n), 空间 O(1)

import (
	"fmt"
)

func main() {
	s := "leetcode"
	fmt.Println(reverseVowels(s))

	s = "hello"
	fmt.Println(reverseVowels(s))
}

func reverseVowels(s string) string {

	vowels := map[byte]byte{
		'a': 'a', 'e': 'e', 'i': 'i', 'o': 'o', 'u': 'u',
		'A': 'A', 'E': 'E', 'I': 'I', 'O': 'O', 'U': 'U',
	}
	res := []byte(s)

	p1 := 0
	p2 := len(s) - 1

	for p1 < p2 {
		v1, ok1 := vowels[s[p1]]
		v2, ok2 := vowels[s[p2]]

		if ok1 && !ok2 {
			p2--
			continue
		}

		if !ok1 && ok2 {
			p1++
			continue
		}

		if ok1 && ok2 {
			res[p1], res[p2] = v2, v1
		}
		p1++
		p2--
	}

	return string(res)
}
