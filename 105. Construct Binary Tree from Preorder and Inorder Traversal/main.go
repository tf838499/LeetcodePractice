package main

import "fmt"

// "fmt"
// "strconv"
// "strings"
// "math"

/*
Given two integer arrays inorder and postorder where inorder
is the inorder traversal of a binary tree and postorder
is the postorder traversal of the same tree,
construct and return the binary tree.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	med := preorder[0]
	ind := 0
	for inorder[ind] != med {
		ind++
	}
	preorder = preorder[1:]
	return &TreeNode{Val: med,
		Left:  buildTree(preorder[0:ind], inorder[0:ind]),
		Right: buildTree(preorder[ind:], inorder[ind+1:]),
	}
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	P := buildTree(preorder, inorder)
	fmt.Println(P)
}
