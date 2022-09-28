package main

import (
// "fmt"
// "strconv"
// "strings"
)

/*
Easy
iteration
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

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}

	stack := make([]*TreeNode, 0)
	// var tmp *TreeNode
	cur := root
	for len(stack) != 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			ans = append(ans, top.Val)
			cur = top.Right
		}
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

	inorderTraversal(Node1)
}
