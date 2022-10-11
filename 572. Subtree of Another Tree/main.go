package main

import (
	"fmt"
	// "strconv"
	// "strings"
	// "math"
)

/*
Easy
recursive
dfs
tree
done

Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

Input: root = [1,2,2,3,4,4,3]
Output: true
Input: root = [1,2,2,null,3,null,3]
Output: false
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root.Left == nil || root.Right == nil {
		return false
	}
	if root.Val == subRoot.Val && (root.Left.Val == subRoot.Left.Val) && (root.Right.Val == subRoot.Right.Val) {

		if root.Left.Left != nil || root.Left.Right != nil {
			return false
		}
		if root.Right.Left != nil || root.Right.Right != nil {
			return false
		}
		return true
	}
	Leftbool := isSubtree(root.Left, subRoot)
	Rightbool := isSubtree(root.Right, subRoot)
	ans := Leftbool || Rightbool
	return ans
}

func isSubtree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isEqual(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isEqual(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && isEqual(a.Left, b.Left) && isEqual(a.Right, b.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}
	if s.Val == t.Val && isIdentical(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isIdentical(s, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}
	return s.Val == t.Val && isIdentical(s.Left, t.Left) && isIdentical(s.Right, t.Right)
}

// [3,4,5,1,2,null,null,null,null,0]
// [4,1,2]
func main() {
	var Node6 *TreeNode = &TreeNode{
		Val: 0.0,
	}
	var Node5 *TreeNode = &TreeNode{
		Val: 2.0, Right: Node6,
	}
	var Node4 *TreeNode = &TreeNode{
		Val: 1.0,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 5.0,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 4.0, Left: Node4, Right: Node5,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 3.0, Left: Node2, Right: Node3,
	}

	var Node3_1 *TreeNode = &TreeNode{
		Val: 2.0,
	}
	var Node2_1 *TreeNode = &TreeNode{
		Val: 1.0,
	}
	var Node1_1 *TreeNode = &TreeNode{
		Val: 4.0, Left: Node2_1, Right: Node3_1,
	}

	a := isSubtree(Node1, Node1_1)
	fmt.Println(a)
}
