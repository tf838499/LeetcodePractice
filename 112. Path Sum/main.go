package main

import (
	"fmt"
	// "strconv"
	// "strings"
	// "math"
)

/*
Easy
recursive
dfs
tree
done

1. Please don't post any solutions in this discussion tab.

2. The problem discussion tab is for asking questions about the problem or for sharing tips - anything except for solutions.

3. If you'd like to share your solution for feedback and ideas, please head to the solutions tab and post it there.

Given the root of a binary tree and an integer targetSum,
return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

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

	divideSum := targetSum - root.Val
	result := false
	if divideSum == 0 && root.Right == nil && root.Left == nil {
		result = true
		return result
	}

	result = hasPathSum(root.Left, divideSum) || hasPathSum(root.Right, divideSum)
	return result
}

func main() {

	var Node1_4 *TreeNode = &TreeNode{
		Val: 2.0,
	}
	var Node1_3 *TreeNode = &TreeNode{
		Val: 7.0,
	}

	// var Node3 *TreeNode = &TreeNode{
	// 	Val: 2.0,
	// }
	// var Node2 *TreeNode = &TreeNode{
	// 	Val: 2.0,
	// }
	var Node1_2 *TreeNode = &TreeNode{
		Val: 11.0, Left: Node1_3, Right: Node1_4,
	}

	var Node3_1 *TreeNode = &TreeNode{
		Val: 8.0,
	}
	var Node2_1 *TreeNode = &TreeNode{
		Val: 4.0, Left: Node1_2,
	}
	var Node1_1 *TreeNode = &TreeNode{
		Val: 5.0, Left: Node2_1, Right: Node3_1,
	}
	targetSum := 22
	a := hasPathSum(Node1_1, targetSum)
	fmt.Println(a)
}
