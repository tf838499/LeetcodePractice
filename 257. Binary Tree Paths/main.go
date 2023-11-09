package main

import "strconv"

// "fmt"

// "strings"
// "math"

/*
easy
dfs
recursion
doing

Given the root of a binary tree,
return all root-to-leaf paths in any order.

A leaf is a node with no children.

Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func binaryTreePaths(root *TreeNode) []string {
// 	// dfs := *[]TreeNode{}
// }
func binaryTreePaths(root *TreeNode) []string {
	tmp := []string{}
	ans := []string{}
	if root == nil {
		return ans
	}

	stack := []*TreeNode{}
	stack = append(stack, root)
	tmp = append(tmp, strconv.Itoa(root.Val))
	for len(stack) != 0 {
		lenStack := len(stack)
		for i := 0; i < lenStack; i++ {
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
				tmp = append(tmp, tmp[i]+"->"+strconv.Itoa(stack[i].Left.Val))
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
				tmp = append(tmp, tmp[i]+"->"+strconv.Itoa(stack[i].Right.Val))
			}
			if stack[i].Left == nil && stack[i].Right == nil {
				ans = append(ans, tmp[i])
			}
		}
		stack = stack[lenStack:]
		tmp = tmp[lenStack:]
	}
	return ans
}

func main() {
	// var Node7 *TreeNode = &TreeNode{
	// 	Val: 7.0,
	// }
	// var Node6 *TreeNode = &TreeNode{
	// 	Val: 6.0,
	// }
	// var Node5 *TreeNode = &TreeNode{
	// 	Val: 5.0,
	// }
	var Node4 *TreeNode = &TreeNode{
		Val: 4.0,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 3.0,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2.0, Left: Node4,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2, Right: Node3,
	}

	binaryTreePaths(Node1)
}

// [1,2,3,4,5,6,null,8]
