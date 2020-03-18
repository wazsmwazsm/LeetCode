// https://leetcode-cn.com/problems/linked-list-cycle/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	l4 := &ListNode{
		Val: 4,
	}
	l3 := &ListNode{
		Val:  0,
		Next: l4,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  3,
		Next: l2,
	}
	l4.Next = l2

	fmt.Println(hasCycle(l1))
}

// 利用 map 存下遍历的节点, 如果已存, 则有环
// 时间 O(n) 空间 O(n)
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = true
		head = head.Next
	}

	return false
}
