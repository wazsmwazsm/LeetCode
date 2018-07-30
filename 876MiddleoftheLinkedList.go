/*
Given a non-empty, singly linked list with head node head, return a middle node of linked list.

If there are two middle nodes, return the second middle node.

Example 1:

Input: [1,2,3,4,5]
Output: Node 3 from this list (Serialization: [3,4,5])
The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).
Note that we returned a ListNode object ans, such that:
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.
Example 2:

Input: [1,2,3,4,5,6]
Output: Node 4 from this list (Serialization: [4,5,6])
Since the list has two middle nodes with values 3 and 4, we return the second one.
 

Note:

The number of nodes in the given list will be between 1 and 100.
*/
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
