package main

import (
	// "fmt"
	"strconv"
	// "strings"
	// "math"
)

/*
easy
dfs
recursion
doing

Given the root of a binary tree, return all root-to-leaf paths in any order.

A leaf is a node with no children.

Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func binaryTreePaths(root *TreeNode) []string{
	ans := []string{}
	travel = func()
} 
// func binaryTreePaths(root *TreeNode) []string {
// 	res := make([]string, 0)
// 	var travel func(node *TreeNode, s string)

// 	travel = func(node *TreeNode, s string) {
// 		if node.Left == nil && node.Right == nil {
// 			v := s + strconv.Itoa(node.Val)
// 			res = append(res, v)
// 			return
// 		}
// 		s = s + strconv.Itoa(node.Val) + "->"
// 		if node.Left != nil {
// 			travel(node.Left, s)
// 		}
// 		if node.Right != nil {
// 			travel(node.Right, s)
// 		}
// 	}

// 	travel(root, "")
// 	return res
// }


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
