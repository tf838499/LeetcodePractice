package main

import (
// "fmt"
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

Given the roots of two binary trees p and q,
write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical,
 and the nodes have the same value.

Input: root = [1,2,2,3,4,4,3]
Output: true
Input: root = [1,2,2,null,3,null,3]
Output: false
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val == q.Val {
		// isSameTree(p.Left,q.Left)
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false

}

func main() {

	var Node3 *TreeNode = &TreeNode{
		Val: 2.0,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2.0,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2, Right: Node3,
	}

	var Node3_1 *TreeNode = &TreeNode{
		Val: 2.0,
	}
	var Node2_1 *TreeNode = &TreeNode{
		Val: 2.0,
	}
	var Node1_1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2_1, Right: Node3_1,
	}

	isSameTree(Node1, Node1_1)
}
