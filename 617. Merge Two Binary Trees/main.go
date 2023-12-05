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

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// if one of t1 and t2 is nil, return the other
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	// merge t1 and t2
	root := &TreeNode{Val: t1.Val + t2.Val}
	// recursion
	root.Left = mergeTrees(t1.Left, t2.Left)
	root.Right = mergeTrees(t1.Right, t2.Right)
	return root
}
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// if root1 == nil && root2 == nil {
	// 	return nil
	// }
	// if root1 != nil && root2 != nil {
	// 	root1.Val = root1.Val + root2.Val
	// }
	if root1 == nil {
		root1 = root2
		return root1
	}
	if root2 == nil {
		return root1
	}
	root := &TreeNode{Val: root1.Val + root2.Val}
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}

func main() {
	// num := []int{3, 2, 1, 6, 0, 5}
	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 5,
	// }
	var Node3 *TreeNode = &TreeNode{
		Val: 3,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2, Left: Node3,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1, Left: Node2, Right: nil,
	}

	// var Node5_1 *TreeNode = &TreeNode{
	// 	Val: 7,
	// }
	// var Node4_1 *TreeNode = &TreeNode{
	// 	Val: 4,
	// }
	var Node3_1 *TreeNode = &TreeNode{
		Val: 3,
	}
	var Node2_1 *TreeNode = &TreeNode{
		Val: 2, Right: Node3_1,
	}
	var Node1_1 *TreeNode = &TreeNode{
		Val: 1, Right: Node2_1,
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
