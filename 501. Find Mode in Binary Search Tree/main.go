package main

import (
// "fmt"
// "strconv"
// "strings"
// "math"
)

/*
Given the root of a binary search tree (BST) with duplicates,
return all the mode(s) (i.e., the most frequently occurred element) in it.

If the tree has more than one mode, return them in any order.

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than or
equal to the node's key.
The right subtree of a node contains only nodes with keys greater
than or equal to the node's key.
Both the left and right subtrees must also be binary search trees.
Input: root = [1,null,2,2]
Output: [2]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func findMode(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{0}
// 	}
// 	recodeMap := make(map[int]int)
// 	ans := findMaxMode(root, recodeMap)
// 	return []int{ans}
// }
// func findMaxMode(root *TreeNode, recordMap map[int]int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	recordMap[root.Val]++
// 	findMaxR := findMaxMode(root.Right, recordMap)
// 	findMaxL := findMaxMode(root.Left, recordMap)
// 	if recordMap[findMaxL] > recordMap[findMaxR] {
// 		if recordMap[findMaxL] > recordMap[root.Val] {
// 			return findMaxL
// 		} else {
// 			return root.Val
// 		}
// 	} else {
// 		if recordMap[findMaxR] > recordMap[root.Val] {
// 			return findMaxR
// 		} else {
// 			return root.Val
// 		}
// 	}
// }
func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	count := 1
	max := 1
	var prev *TreeNode
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && prev.Val == node.Val {
			count++
		} else {
			count = 1
		}
		if count >= max {
			if count > max && len(res) > 0 {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			max = count
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return res
}
func main() {

	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 4.0,
	// }
	// var Node5 *TreeNode = &TreeNode{
	// 	Val: 4,
	// }
	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 3,
	// }
	var Node3 *TreeNode = &TreeNode{
		Val: 2,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2, Right: Node3,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1, Left: Node2,
	}
	findMode(Node1)
}
