package main

// "fmt"
// "strconv"
// "strings"
// "math"

/*
Easy
recursive
dfs
tree
done

Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	ans := compare(root.Left, root.Right)

	return ans
	// return
}
func compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil {
		return false
	} else if right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}

func main() {
	var Node7 *TreeNode = &TreeNode{
		Val: 3.0,
	}
	// var Node6 *TreeNode = &TreeNode{
	// 	Val: 4.0,
	// }
	var Node5 *TreeNode = &TreeNode{
		Val: 4.0,
	}
	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 3.0,
	// }
	var Node3 *TreeNode = &TreeNode{
		Val: 2.0, Right: Node7,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2.0, Right: Node5,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2, Right: Node3,
	}

	isSymmetric(Node1)
}
