package main

import (
	"fmt"
	// "strconv"
)

/*
You are given two binary trees root1 and root2.

Imagine that when you put one of them to cover the other,
some nodes of the two trees are overlapped while the others are not.
You need to merge the two trees into a new binary tree.
The merge rule is that if two nodes overlap,
then sum node values up as the new value of the merged node. Otherwise,
the NOT null node will be used as the node of the new tree.

Return the merged tree.

Note: The merging process must start from the root nodes of both trees.
*/

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// if root1 == nil && root2 == nil {
	// 	return nil
	// } else if root1 == nil {
	// 	return root2
	// 	// return root1
	// } else if root2 == nil {
	// 	return root1
	// }
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val = root1.Val + root2.Val
	root1.Right = mergeTrees(root1.Right, root2.Right)
	root1.Left = mergeTrees(root1.Left, root2.Left)

	return root1
}

func main() {
	// num := []int{3, 2, 1, 6, 0, 5}
	var Node4 *TreeNode = &TreeNode{
		Val: 5.0,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 3.0, Left: Node4,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2.0,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node3, Right: Node2,
	}

	var Node5_1 *TreeNode = &TreeNode{
		Val: 7.0,
	}
	var Node4_1 *TreeNode = &TreeNode{
		Val: 4.0,
	}
	var Node3_1 *TreeNode = &TreeNode{
		Val: 3.0, Right: Node5_1,
	}
	var Node2_1 *TreeNode = &TreeNode{
		Val: 1.0, Right: Node4_1,
	}
	var Node1_1 *TreeNode = &TreeNode{
		Val: 2.0, Left: Node2_1, Right: Node3_1,
	}
	p := mergeTrees(Node1, Node1_1)
	fmt.Println(p)
	// node5 := &TreeNode{Val: 3}
	// node4 := &TreeNode{Val: 1}
	// node3 := &TreeNode{Val: 7}
	// node2 := &TreeNode{Val: 2, Left: node4, Right: node5}
	// node1 := &TreeNode{Val: 4, Left: node2, Right: node3}

	// num := []int{-1, 0, 3, 5, 9, 12} // 0
	// target := 9

}
