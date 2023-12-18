package main

import (
	"fmt"
	"math"
	// "strconv"
)

/*
Given the root of a Binary Search Tree (BST),
return the minimum absolute difference between
the values of any two different nodes in the tree.
*/

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	x = math.MaxInt
	minVal := math.MaxInt
	minVal = inorder(root, minVal)
	return minVal
}

var x int

func inorder(root *TreeNode, minVal int) int {
	if root == nil {
		return 0
	}
	ans := inorder(root.Left, minVal)
	if ans != 0 {
		minVal = ans
	}
	p := x - root.Val
	if p < 0 {
		p = -p
	}
	if p < minVal {
		minVal = p
	}
	x = root.Val
	ans = inorder(root.Right, minVal)
	if ans != 0 {
		minVal = ans
	}
	return minVal
}

// var MinDiff, X int

// func getMinimumDifference(root *TreeNode) int {
//     X = math.MaxInt64
//     MinDiff = math.MaxInt64
//     inorder(root)
//     return MinDiff
// }

// func inorder(root *TreeNode) {
//     if root == nil {
//         return
//     }
//     inorder(root.Left)

//     if root.Val > X {
//         MinDiff = min(MinDiff, root.Val-X)
//     } else {
//         MinDiff = min(MinDiff, X-root.Val)
//     }
//     X = root.Val

//	    inorder(root.Right)
//	    return
//	}
func main() {

	var Node5 *TreeNode = &TreeNode{
		Val: 911,
	}
	var Node4 *TreeNode = &TreeNode{
		Val: 227,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 701, Left: nil, Right: Node5,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 104, Left: nil, Right: Node4,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 236, Left: Node2, Right: Node3,
	}
	// 		236
	//	104		701
	//0   277  0   911
	// var Node5 *TreeNode = &TreeNode{
	// 	Val: 49,
	// }
	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 12,
	// }
	// var Node3 *TreeNode = &TreeNode{
	// 	Val: 48, Left: Node4, Right: Node5,
	// }
	// var Node2 *TreeNode = &TreeNode{
	// 	Val: 0,
	// }
	// var Node1 *TreeNode = &TreeNode{
	// 	Val: 1, Left: Node2, Right: Node3,
	// }

	p := getMinimumDifference(Node1)
	fmt.Println(p)
}
