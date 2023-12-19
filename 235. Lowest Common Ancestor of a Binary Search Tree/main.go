package main

// "fmt"
// "strconv"
// "strings"
// "math"

/*
Given a binary search tree (BST),
find the lowest common ancestor (LCA) node
of two given nodes in the BST.

According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between
two nodes p and q as the lowest node in T
that has both p and q as descendants
(where we allow a node to be a descendant of itself).”

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2,
since a node can be a descendant of itself according to the LCA definition.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root == nil {
// 		return root
// 	}
// 	if root == p || root == q {
// 		return root
// 	}
// 	var U, V *TreeNode
// 	if root.Val > p.Val && root.Val > q.Val {
// 		U = lowestCommonAncestor(root.Left, p, q)
// 	} else if root.Val < p.Val && root.Val < q.Val {
// 		V = lowestCommonAncestor(root.Right, p, q)
// 	} else {
// 		U = lowestCommonAncestor(root.Left, p, q)
// 		V = lowestCommonAncestor(root.Right, p, q)
// 	}
// 	if U != nil && V != nil {
// 		return root
// 	}
// 	if U != nil {
// 		return U
// 	} else {
// 		return V
// 	}
// 	// if

// 	// return nil
// }
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
func main() {

	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 4.0,
	// }
	var Node9 *TreeNode = &TreeNode{
		Val: 9,
	}
	var Node8 *TreeNode = &TreeNode{
		Val: 7,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 8, Left: Node8, Right: Node9,
	}
	var Node7 *TreeNode = &TreeNode{
		Val: 5,
	}
	var Node6 *TreeNode = &TreeNode{
		Val: 3,
	}
	var Node5 *TreeNode = &TreeNode{
		Val: 4, Left: Node6, Right: Node7,
	}
	var Node4 *TreeNode = &TreeNode{
		Val: 0,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2, Left: Node4, Right: Node5,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 6, Left: Node2, Right: Node3,
	}
	c := lowestCommonAncestor(Node1, Node2, Node5)
	print(c)
}
