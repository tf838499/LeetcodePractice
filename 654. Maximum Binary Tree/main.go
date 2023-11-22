package main

import (
	"fmt"
	// "strconv"
)

/*
You are given an integer array nums with no duplicates.
A maximum binary tree can be built recursively from nums using the following algorithm:

Create a root node whose value is the maximum value in nums.
Recursively build the left subtree on the subarray prefix to the left of the maximum value.
Recursively build the right subtree on the subarray suffix to the right of the maximum value.
Return the maximum binary tree built from nums.
Input: nums = [3,2,1,6,0,5]
Output: [6,3,5,null,2,0,null,null,1]
Input: nums = [3,2,1]
Output: [3,null,2,null,1]
*/

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[index] < nums[i] {
			index = i
		}
	}
	left := constructMaximumBinaryTree(nums[0:index])
	right := constructMaximumBinaryTree(nums[index+1:])
	return &TreeNode{Val: nums[index], Right: right, Left: left}
}

func main() {
	num := []int{3, 2, 1, 6, 0, 5}
	p := constructMaximumBinaryTree(num)
	fmt.Println(p)

}
