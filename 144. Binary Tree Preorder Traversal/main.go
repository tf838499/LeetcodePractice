package main

import "fmt"

// "fmt"
// "strconv"
// "strings"

/*
Easy
recursive
binary tree
done

Given the root of a binary tree, return the preorder traversal of its nodes' values.

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

// func preorderTraversal(root *TreeNode) []int {
// 	ans := []int{}
// 	if root != nil {
// 		ans = append(ans, root.Val)
// 		ans = append(ans, preorderTraversal(root.Left)...)
// 		ans = append(ans, preorderTraversal(root.Right)...)
// 		// preorderTraversal(root.Right)
// 	}
// 	return ans
// }
// func preorderTraversal(root *TreeNode) []int {
// 	ans := []int{}
// 	if root == nil {
// 		return ans
// 	}
// 	ans = append(ans, root.Val)
// 	ans = append(ans, preorderTraversal(root.Left)...)
// 	ans = append(ans, preorderTraversal(root.Right)...)
// 	return ans
// }
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		ans = append(ans, root.Val)
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		root = queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]
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
	p := preorderTraversal(Node1)
	fmt.Println(p)
}
