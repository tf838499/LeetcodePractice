package main

import (
	"fmt"
	// "strconv"
	// "strings"
	// "math"
)

/*


Given the root of a binary tree and an integer targetSum,
return true if the tree has a root-to-leaf path such
that adding up all the values along the path equals targetSum.

A leaf is a node with no children.

Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
Input: root = [1,2,3], targetSum = 5
Output: false
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {

	// var Node1_4 *TreeNode = &TreeNode{
	// 	Val: 2,
	// }
	// var Node1_3 *TreeNode = &TreeNode{
	// 	Val: 7,
	// }

	// var Node3 *TreeNode = &TreeNode{
	// 	Val: 3,
	// }
	// var Node2 *TreeNode = &TreeNode{
	// 	Val: 2,
	// }
	// var Node1 *TreeNode = &TreeNode{
	// 	Val: 1, Left: Node2, Right: Node3,
	// }
	// var Node1_2 *TreeNode = &TreeNode{
	// 	Val: 11, Left: Node1_3, Right: Node1_4,
	// }

	// var Node3_1 *TreeNode = &TreeNode{
	// 	Val: 8,
	// }
	// var Node2_1 *TreeNode = &TreeNode{
	// 	Val: 4, Left: Node1_2,
	// }
	// var Node1_1 *TreeNode = &TreeNode{
	// 	Val: 5, Left: Node2_1, Right: Node3_1,
	// }
	var NodeNone *TreeNode = &TreeNode{}
	targetSum := 0
	a := hasPathSum(NodeNone, targetSum)
	fmt.Println(a)
}
