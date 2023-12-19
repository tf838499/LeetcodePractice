package main

// "fmt"
// "strconv"
// "strings"
// "math"

/*
Given a binary tree, find the lowest common ancestor (LCA)
 of two given nodes in the tree.

According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two
 nodes p and q as the lowest node in T that has
  both p and q as descendants
  (where we allow a node to be a descendant of itself).
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
	var U, V *TreeNode
	if root.Left != nil {
		U = lowestCommonAncestor(root.Left, p, q)
	}
	if root.Right != nil {
		V = lowestCommonAncestor(root.Right, p, q)
	}
	if U != nil && V != nil {
		return root
	}
	if U != nil {
		return U
	}
	if V != nil {
		return V
	}
	if root == p {
		return root
	}
	if root == q {
		return root
	}

	return nil
}

// 		3
// 	4	 	5
// 6 7    9  8
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
	c := lowestCommonAncestor(Node1, Node2, Node7)
	print(c)
}
