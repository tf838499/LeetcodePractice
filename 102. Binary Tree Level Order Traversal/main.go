package main

import (
// "fmt"
// "strconv"
// "strings"
)

/*
Medium
bfs
done

Given the root of a binary tree, return the level order traversal of its nodes' values.
(i.e., from left to right, level by level).

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	// tmp := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for root != nil && len(stack) != 0 {
		stackLen := len(stack)
		tmp := make([]int, 0)
		for i := 0; i < stackLen; i++ {
			top := stack[i]
			if top == nil {
				continue
			}
			tmp = append(tmp, top.Val)
			stack = append(stack, top.Left)
			stack = append(stack, top.Right)
		}
		stack = stack[stackLen:]
		if len(tmp) != 0 {
			ans = append(ans, tmp)
		}

	}
	return ans

}

func main() {
	var Node5 *TreeNode = &TreeNode{
		Val: 7.0,
	}
	var Node4 *TreeNode = &TreeNode{
		Val: 15.0,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 20.0, Right: Node5, Left: Node4,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 9.0,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 3.0, Right: Node3, Left: Node2,
	}

	levelOrder(Node1)
}
