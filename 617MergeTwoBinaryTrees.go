// see https://leetcode.com/problems/merge-two-binary-trees/description/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val

	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}

func main() {
	a4 := &TreeNode{5, nil, nil}
	a3 := &TreeNode{2, nil, nil}
	a2 := &TreeNode{3, a4, nil}
	a1 := &TreeNode{1, a2, a3}

	b5 := &TreeNode{7, nil, nil}
	b4 := &TreeNode{4, nil, nil}
	b3 := &TreeNode{3, nil, b5}
	b2 := &TreeNode{1, nil, b4}
	b1 := &TreeNode{2, b2, b3}

	fmt.Println(a1.Val, a1.Left.Val, a1.Right.Val, a1.Left.Left.Val)

	merge := mergeTrees(a1, b1)
	fmt.Println(merge.Val, merge.Left.Val, merge.Right.Val, merge.Left.Left.Val, merge.Left.Right.Val, merge.Right.Left, merge.Right.Right.Val)
}
