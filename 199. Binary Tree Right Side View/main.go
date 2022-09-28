package main

import (
// "fmt"
// "strconv"
// "strings"
// "math"
)

/*
Medium
bfs
dfs
queue
done

Given the root of a binary tree,
imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.

Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) (ans []int) {
	stack := []*TreeNode{root}
	if root == nil {
		return
	}
	// ans = append(ans, root.Val)
	for len(stack) != 0 {
		// ans = append(ans, root.Val)
		tmp := []int{}
		for _, i := range stack {
			stack = stack[1:]
			if i.Left != nil {
				stack = append(stack, i.Left)
			}
			if i.Right != nil {
				stack = append(stack, i.Right)
			}
			tmp = append(tmp, i.Val)
		}
		ans = append(ans, tmp[len(tmp)-1])
	}
	return
}

// func rightSideView(root *TreeNode) []int {
// 	var result []int
// 	populateRightSideView(root, &result, 0)
// 	return result
// }

// func populateRightSideView(root *TreeNode, result *[]int, level int) {
// 	if root == nil {
// 		return
// 	}
// 	if level == len(*result) {
// 		*result = append(*result, root.Val)
// 	}
// 	populateRightSideView(root.Right, result, level+1)
// 	populateRightSideView(root.Left, result, level+1)
// }
func main() {

	var Node4 *TreeNode = &TreeNode{
		Val: 4.0,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 3.0,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2.0, Right: Node4,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2, Right: Node3,
	}

	rightSideView(Node1)
}
