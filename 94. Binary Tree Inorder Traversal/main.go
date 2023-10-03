package main

// "strconv"
// "strings"

/*
Easy
recursion
binary tree
done

Given the root of a binary tree, return the inorder traversal of its nodes' values.
L0 → L1 → … → Ln - 1 → Ln
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
*/
/**
 * Definition for singly-linked list.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func inorderTraversal(root *TreeNode) []int {
// 	ans := []int{}
// 	if root != nil {
// 		ans = append(ans, inorderTraversal(root.Left)...)
// 		ans = append(ans, root.Val)
// 		// ans = append(ans, preorderTraversal(root.Left)...)
// 		ans = append(ans, inorderTraversal(root.Right)...)
// 		// preorderTraversal(root.Right)
// 	}
// 	return ans
// }
// func inorderTraversal(root *TreeNode) []int {
// 	ans := []int{}
// 	if root != nil {
// 		ans = append(ans, inorderTraversal(root.Left)...)
// 		ans = append(ans, root.Val)
// 		ans = append(ans, inorderTraversal(root.Right)...)
// 	}
// 	return ans
// }
// func inorderTraversal(root *TreeNode) []int {
// 	ans := []int{}
// 	tmp := []*TreeNode{}

// 	if root.Left == nil {
// 		ans = append(ans, root.Val)
// 		tmp = tmp[:len(tmp)-1]
// 		tmp = append(tmp, root.Right)
// 	}
// 	if root.Left != nil {
// 		tmp = append(tmp, root)
// 		root = root.Left
// 	}

// 	return ans
// }
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
func main() {

	var Node3 *TreeNode = &TreeNode{
		Val: 3.0,
	}

	var Node2 *TreeNode = &TreeNode{
		Val: 2.0,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2, Right: Node3,
	}
	inorderTraversal(Node1)
}
