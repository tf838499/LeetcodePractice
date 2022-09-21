package main

import (
// "fmt"
// "strconv"
// "strings"
// "math"
)

/*
Given the root of a complete binary tree, return the number of the nodes in the tree.

According to Wikipedia, every level, except possibly the last,
is completely filled in a complete binary tree,
and all nodes in the last level are as far left as possible.
It can have between 1 and 2h nodes inclusive at the last level h.

Design an algorithm that runs in less than O(n) time complexity.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}
	ans = 1 + countNodes(root.Right) + countNodes(root.Left)
	return ans
}

func main() {
	var Node7 *TreeNode = &TreeNode{
		Val: 7.0,
	}
	var Node6 *TreeNode = &TreeNode{
		Val: 6.0,
	}
	var Node5 *TreeNode = &TreeNode{
		Val: 5.0,
	}
	var Node4 *TreeNode = &TreeNode{
		Val: 4.0,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 3.0, Left: Node6, Right: Node7,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2.0, Left: Node4, Right: Node5,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2, Right: Node3,
	}

	countNodes(Node1)
}
