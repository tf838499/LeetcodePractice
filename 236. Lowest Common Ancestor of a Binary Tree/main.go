package main

import (
// "fmt"
// "strconv"
// "strings"
// "math"
)

/*
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	var U, I *TreeNode

	I = lowestCommonAncestor(root.Left, p, q)
	U = lowestCommonAncestor(root.Right, p, q)

	if I != nil && U != nil {
		return root
	}

	if I != nil {
		return I
	}
	if U != nil {
		return U
	}
	return nil
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	// check
// 	if root == nil {
// 		return root
// 	}
// 	// 相等 直接返回root节点即可
// 	if root == p || root == q {
// 		return root
// 	}
// 	// Divide
// 	left := lowestCommonAncestor(root.Left, p, q)
// 	right := lowestCommonAncestor(root.Right, p, q)

// 	// Conquer
// 	// 左右两边都不为空，则根节点为祖先
// 	if left != nil && right != nil {
// 		return root
// 	}
// 	if left != nil {
// 		return left
// 	}
// 	if right != nil {
// 		return right
// 	}
// 	return nil
// }

// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// Output: 5
// Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.
func main() {

	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 4.0,
	// }
	var Node9 *TreeNode = &TreeNode{
		Val: 8,
	}
	var Node8 *TreeNode = &TreeNode{
		Val: 0,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 1, Left: Node8, Right: Node9,
	}
	var Node7 *TreeNode = &TreeNode{
		Val: 4,
	}
	var Node6 *TreeNode = &TreeNode{
		Val: 7,
	}
	var Node5 *TreeNode = &TreeNode{
		Val: 2, Left: Node6, Right: Node7,
	}
	var Node4 *TreeNode = &TreeNode{
		Val: 6,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 5, Left: Node4, Right: Node5,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 3, Left: Node2, Right: Node3,
	}
	c := lowestCommonAncestor(Node1, Node2, Node9)
	print(c)
}
