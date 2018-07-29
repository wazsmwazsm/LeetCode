// see: https://leetcode.com/problems/middle-of-the-linked-list/description/
package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    
    fast := head
    slow := head
    
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    
    return slow
}

func main() {
    l6 := &ListNode{6, nil}
    l5 := &ListNode{5, l6}
    l4 := &ListNode{4, l5}
    l3 := &ListNode{3, l4}
    l2 := &ListNode{2, l3}
    l1 := &ListNode{1, l2}
    
    list := middleNode(l1)
    
    for list != nil {
        fmt.Println(list.Val)
        list = list.Next
    }
}
