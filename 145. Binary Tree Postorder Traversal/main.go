package main

// "strconv"
// "strings"

/*
Easy
recursive
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

// func postorderTraversal(root *TreeNode) []int {
// 	ans := []int{}
// 	if root != nil {
// 		ans = append(ans, postorderTraversal(root.Left)...)
// 		ans = append(ans, postorderTraversal(root.Right)...)
// 		ans = append(ans, root.Val)
// 	}
// 	return ans
// }
// func postorderTraversal(root *TreeNode) []int {
// 	ans := []int{}
// 	if root != nil {
// 		ans = append(ans, postorderTraversal(root.Left)...)
// 		ans = append(ans, postorderTraversal(root.Right)...)
// 		ans = append(ans, root.Val)
// 	}
// 	return ans
// }
func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {

		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		ans = append(ans, root.Val)
		root = queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]

		// root = queue[len(queue)-1]
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}
func main() {
	var Node3 *TreeNode = &TreeNode{
		Val: 3.0,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2.0, Left: Node3,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Right: Node2,
	}
	postorderTraversal(Node1)
}
