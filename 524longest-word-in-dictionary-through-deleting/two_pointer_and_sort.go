// https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "abpcplea"
	d := []string{"ale", "apple", "monkey", "plea"}
	// d := []string{"ba", "ab", "monkey", "ss"}

	word := findLongestWord(s, d)
	fmt.Println(word)
}

// 先给词组排序, 再遍历找子串
// 时间：排序 O(logn) (快排), 遍历 O(n) and 子串 O(n) = O(n^2)
// 时间复杂度为 O(n^2) + O(logn) = O(n^2)
// 空间：子串查找需要创建变量, 遍历后空间为 O(n), 排序(快排)空间为 O(logn)
// 空间复杂度为 O(logn)

func findLongestWord(s string, d []string) string {
	var word string
	var maxLen int

	sort.Strings(d)

	for _, v := range d {
		vLen := len(v)
		if vLen == maxLen {
			continue
		}

		if isSub(s, v) && vLen >= maxLen {
			word = v
			maxLen = vLen
		}
	}

	return word
}

// 双指针去匹配，看总匹配的个数和 word 是不是相等
func isSub(s, d string) bool {
	var p1 int
	var p2 int

	for p1 < len(s) && p2 < len(d) {
		if s[p1] == d[p2] {
			p2++
		}
		p1++
	}

	return p2 == len(d)
}
