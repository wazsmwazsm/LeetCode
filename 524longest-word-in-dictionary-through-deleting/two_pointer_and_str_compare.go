// https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abpcplea"
	d := []string{"ale", "apple", "monkey", "plea"}
	// d := []string{"ba", "ab", "monkey", "ss"}

	word := findLongestWord(s, d)
	fmt.Println(word)
}

// 遍历词组挨个匹配，取最长词，如果长度相等，
// 取字符串排行小的 (直接 string compare (底层字符数组 []byte 对比。其实就是比每个字符 int 值大小))
// 时间：是否是子串 O(n), 遍历为 O(n), （strings.Compare 算的话其实也是 O(n)，但和子串检查是 + 的关系，算在一起）
// 综上时间复杂度为 O(n^2)
// 空间 O(n) （判断字串时会创建变量）
func findLongestWord(s string, d []string) string {
	var word string
	var maxLen int

	for _, v := range d {
		vLen := len(v)
		if vLen == maxLen && strings.Compare(word, v) == -1 {
			// if vLen == maxLen && word < v { // 这样也可以（不确定其它语言）
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
