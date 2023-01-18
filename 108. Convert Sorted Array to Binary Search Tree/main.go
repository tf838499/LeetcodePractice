package main

import (
	"fmt"
	// "strconv"
)

/*
Given an integer array nums where the elements are sorted in ascending order, convert it to a
height-balanced
binary search tree
*/

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	// var ans *TreeNode
	if len(nums) == 0 {
		return ans
	}
	// tmp := ans
	arrayLen := len(nums) / 2
	// if tmp == nil {
	// 	tmp = &TreeNode{Val: nums[arrayLen]}
	// }
	tmp := &TreeNode{Val: nums[arrayLen]}
	tmp.Left = sortedArrayToBST(nums[:arrayLen])
	tmp.Right = sortedArrayToBST(nums[arrayLen+1:])
	// if len(nums) == 1 {
	// 	tmp.Right = sortedArrayToBST(nums[:arrayLen])
	// } else {
	// 	tmp.Left = sortedArrayToBST(nums[:arrayLen])
	// 	tmp.Right = sortedArrayToBST(nums[arrayLen+1:])
	// }

	return tmp
}

func main() {

	// node5 := &TreeNode{Val: 1}
	node4 := []int{-10, -8, -3, -2, 0, 5, 9}

	// [50,30,70,null,40,60,80]
	p := sortedArrayToBST(node4)
	fmt.Println(p)
}
