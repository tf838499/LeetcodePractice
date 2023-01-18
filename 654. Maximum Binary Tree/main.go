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
	ans := &TreeNode{}
	index := findMax(nums)
	ans.Val = nums[index]

	ans.Left = constructMaximumBinaryTree(nums[:index])
	ans.Right = constructMaximumBinaryTree(nums[index+1:])
	return ans
}
func findMax(nums []int) int {
	l, r := 0, len(nums)-1
	fmt.Println(nums)
	for l < r {

		if nums[l] > nums[r] {
			r--
		} else {
			l++
		}
	}
	return l
}

func main() {
	num := []int{3, 2, 1, 6, 0, 5}
	p := constructMaximumBinaryTree(num)
	fmt.Println(p)
	// node5 := &TreeNode{Val: 3}
	// node4 := &TreeNode{Val: 1}
	// node3 := &TreeNode{Val: 7}
	// node2 := &TreeNode{Val: 2, Left: node4, Right: node5}
	// node1 := &TreeNode{Val: 4, Left: node2, Right: node3}

	// num := []int{-1, 0, 3, 5, 9, 12} // 0
	// target := 9

}
