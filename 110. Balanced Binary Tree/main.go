package main

// "fmt"
// "strconv"
// "strings"
// "math"

/*
Easy
recursion
binary tree
done

Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	res, _ := dfs(root)
	return res
}
func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftBool, leftHeight := dfs(root.Left)
	rightBool, rightHeight := dfs(root.Right)
	diff := abs(leftHeight - rightHeight)
	// depth := 1 + max(leftHeight, rightHeight)
	if leftBool && rightBool && diff <= 1 {
		return true, 1 + max(leftHeight, rightHeight)
	}
	return false, -1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 	1
//   2	  3
//  4 5  6
//  8
// [1,2,3,4,5,6,null,8]

func main() {
	var Node8 *TreeNode = &TreeNode{
		Val: 7.0,
	}
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
		Val: 4.0, Left: Node8,
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
	// [1,null,2,null,3]
	isBalanced(Node1)
}

// [1,2,3,4,5,6,null,8]
