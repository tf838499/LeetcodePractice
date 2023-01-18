package main

import (
	"fmt"
	// "strconv"
)

/*
Given the root of a binary search tree and the lowest and
highest boundaries as low and high,
trim the tree so that all its elements lies in [low, high].
Trimming the tree should not change the relative structure of
the elements that will remain in the tree
(i.e., any node's descendant should remain a descendant).
It can be proven that there is a unique answer.

Return the root of the trimmed binary search tree.
Note that the root may change depending on the given bounds.
*/

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	} else if root.Val < low {
		return trimBST(root.Right, low, high)
	} else {
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
		return root
	}
}

// func trimBST(root *TreeNode, L int, R int) *TreeNode {
//     return walk(root, L, R)
// }

// func walk(node *TreeNode, L, R int) *TreeNode {
//     if node == nil {
//         return nil
//     }
//     if node.Val < L {
//         return walk(node.Right, L, R)
//     } else if node.Val > R {
//         return walk(node.Left, L, R)
//     } else {
//         node.Left = walk(node.Left, L, R)
//         node.Right = walk(node.Right, L, R)
//         return node
//     }
// }
func main() {

	// node5 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2, Right: node4}
	node1 := &TreeNode{Val: 3, Left: node2, Right: node3}
	// [50,30,70,null,40,60,80]
	p := trimBST(node1, 1, 1)
	fmt.Println(p)
}
