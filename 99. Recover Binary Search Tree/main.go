package main

import "fmt"

// "strconv"
// "strings"

/*
You are given the root of a binary search tree (BST),
where the values of exactly two nodes of
the tree were swapped by mistake.
Recover the tree without changing its structure.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var prev, first, second *TreeNode
	var travel func(*TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil {
			if first == nil && prev.Val > node.Val {
				first = prev
			}
			if first != nil && prev.Val > node.Val {
				second = node
			}
		}
		prev = node
		travel(node.Right)

	}
	travel(root)
	first.Val, second.Val = second.Val, first.Val
}

// 		2 p
// 1    	4 321
// 	   	 3r
// 	1         1
// 3p       2p
// 	2r		  3r
func recoverTree(root *TreeNode) {
	var prev, first, second *TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil {
			if first == nil && prev.Val >= node.Val {
				first = prev
			}
			if first != nil && prev.Val >= node.Val {
				second = node
			}
		}
		prev = node
		dfs(node.Right)
	}
	dfs(root)

	first.Val, second.Val = second.Val, first.Val
}
func main() {

	var Node3 *TreeNode = &TreeNode{
		Val: 2,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 3, Right: Node3,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1, Left: Node2,
	}

	recoverTree(Node1)
	fmt.Println(Node1)
	// fmt.Println(p)
}

// [32,26,47,19,null,null,56,null,27]
// 	 	32
//   26	   47
// 19  n  n  56
//n 27
