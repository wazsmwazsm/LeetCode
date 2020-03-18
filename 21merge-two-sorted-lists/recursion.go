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

// 递归的思想, 且链表有序. 两个列表一次往后走, 对比值谁小谁往后一步 (将链条连过去)
// 用递归的栈保存了后面连接好的链条, 然后回到最前则是已经排好序的链表
/*
	l1: 1   2   4
	l2: 1   3   4

	初次对比, 递归最上层返回 l2 的 1, l1 前进一步. 二次对比, 2 大于 1, l2 前进一步, 递归二层的结果 (返回 l1)
	连接到最上层的返回结果 l2 的 Next 上 (l1 的 2 节点连接了 l2 的 1 节点). 以此类推, 最后当其中一个链表走完
	上一层返回值的 next 必定连接的是另一个链表的值

	时间 O(n) (n+m)
	空间 O(n) (n+m)
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
