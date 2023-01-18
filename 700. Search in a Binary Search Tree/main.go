package main

import (
	"fmt"
	// "strconv"
)

/*
You are given the root of a binary search tree (BST) and an integer val.

Find the node in the BST
that the node's value equals val and return the subtree rooted with that node.
If such a node does not exist, return null.

*/

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
	return root
}
func main() {
	// nput: root = [4,2,7,1,3], val = 5
	// Output: []
	// [4,2,7,1,3], val = 2
	// tput: [2,1,3]
	node5 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 7}
	node2 := &TreeNode{Val: 2, Left: node4, Right: node5}
	node1 := &TreeNode{Val: 4, Left: node2, Right: node3}

	// num := []int{-1, 0, 3, 5, 9, 12} // 0
	// target := 9
	p := searchBST(node1, 5)
	fmt.Println(p)
}
