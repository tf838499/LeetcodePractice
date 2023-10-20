package main

// "fmt"
// "strconv"
// "strings"
// "math"

/*
easy
recursion
dfs
done


Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
/* recursive */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
func main() {
	var Node7 *TreeNode = &TreeNode{
		Val: 3.0,
	}
	// var Node6 *TreeNode = &TreeNode{
	// 	Val: 4.0,
	// }
	var Node5 *TreeNode = &TreeNode{
		Val: 4.0,
	}
	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 3.0,
	// }
	var Node3 *TreeNode = &TreeNode{
		Val: 2.0, Right: Node7,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2.0, Right: Node5,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2, Right: Node3,
	}

	maxDepth(Node1)
}
