package main

import "fmt"

/**

  LeetCode

  Invert a binary tree.
       4
     /   \
    2     7
   / \   / \
  1   3 6   9
  to
       4
     /   \
    7     2
   / \   / \
  9   6 3   1


*/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

    if (root == nil) {
        return nil
    }

    root.Left, root.Right = root.Right, root.Left

    invertTree(root.Left)
    invertTree(root.Right)

    return root

}

func main() {

    l2 := &TreeNode{3, nil, nil};
    l3 := &TreeNode{4, nil, nil};
    r2 := &TreeNode{6, nil, nil};
    r3 := &TreeNode{7, nil, nil};
    l1 := &TreeNode{2, l2, l3};
    r1 := &TreeNode{5, r2, r3};
    tree := &TreeNode{1, l1, r1};

    fmt.Printf("%v%v%v%v%v%v%v\n", tree.Val, tree.Left.Val, tree.Left.Left.Val, tree.Left.Right.Val, tree.Right.Val, tree.Right.Left.Val, tree.Right.Right.Val);
    // 执行转换
    invertTree(tree);

    fmt.Printf("%v%v%v%v%v%v%v\n", tree.Val, tree.Left.Val, tree.Left.Left.Val, tree.Left.Right.Val, tree.Right.Val, tree.Right.Left.Val, tree.Right.Right.Val);

}
