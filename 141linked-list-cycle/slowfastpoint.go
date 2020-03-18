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

// 利用快慢指针, 一块一慢, 速度差为 1, 快指针会在一定时间相遇
// 时间, 以慢指针为标准计算。未入环距离未 n, 入环后, 和快指针的距离为 m，快慢指针速度差为 1
// 那么需要 1 * k 的时间快指针会追上慢指针。
// 综上, 时间复杂度为 O(n) (n+m)
// 空间 O(1)
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
