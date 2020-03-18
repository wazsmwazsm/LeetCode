// https://leetcode-cn.com/problems/merge-two-sorted-lists/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := ""
	for l != nil {
		s += fmt.Sprintf(" %d ", l.Val)
		l = l.Next
	}

	return s
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	fmt.Printf("l1:%s, l2:%s\n", l1, l2)

	l := mergeTwoLists(l1, l2)

	fmt.Printf("erged: %s", l)
}

// 用一个指针保存链头, 一个游标来调整指向, 根据有序的特点, 创建连接
/*
	l1: 1   2   4
	l2: 1   3   4

	创建链表头，游标指向头
	第一轮对比元素, l1 l2 第一个元素相等, l1 向后一步, 游标指向(next 指向, 并向后一步) l1 的第一个元素
	二轮对比元素, l1 的第二个元素大于 l2 第一个元素, l2 向后一步, 游标指向 l2 的第一个元素
	三轮对比, l1 的第二个元素小于 l2 的第二个元素, l1 向后一步, 游标指向 l2 的第一个元素
	依次类推直到有链表到头, 将未到头的链表剩余部分连接到游标的 next
	返回链表头的 next (因为指针操作, head 指针已经拥有合并后的联调)

	时间 O(n)
	空间 O(1)
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{} // 创建链表头
	prev := head        // 游标指针

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}

		prev = prev.Next
	}

	// 将剩余的元素连接
	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}

	return head.Next
}
