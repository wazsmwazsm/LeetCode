// see https://leetcode.com/problems/binary-tree-pruning/description/ 
package main

import "fmt"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    root.Left = pruneTree(root.Left)
    root.Right = pruneTree(root.Right)
    
    if root.Val == 0 && root.Left == nil && root.Right == nil {
        root = nil
    } 
    
    return root
}


func main() {
    b4 := &TreeNode{1, nil, nil}
    b3 := &TreeNode{0, nil, nil}
    b2 := &TreeNode{1, b3, b4}
    b1 := &TreeNode{1, nil, b2}
	fmt.Println(b1.Val, b1.Left, b1.Right.Val, b1.Right.Left.Val, b1.Right.Right.Val) 
	
    root := pruneTree(b1)
    fmt.Println(root.Val, root.Left, root.Right.Val, root.Right.Left, root.Right.Right.Val) 
}
