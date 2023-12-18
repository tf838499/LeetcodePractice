package main

// "fmt"
// "strconv"
// "strings"
// "math"

/*
Given the root of a binary search tree (BST) with duplicates,
return all the mode(s) (i.e., the most frequently occurred element)
in it.

If the tree has more than one mode, return them in any order.

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than or
equal to the node's key.
The right subtree of a node contains only nodes with keys greater
than or equal to the node's key.
Both the left and right subtrees must also be binary search trees.
Input: root = [1,null,2,2]
Output: [2]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var prevValue *TreeNode
	result := []int{}
	maxValue := -1
	CurrentCount := 0
	var travel func(*TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		if prevValue == nil {
			prevValue = root
		}
		if root.Val == prevValue.Val {
			CurrentCount++
		} else {
			prevValue = root
			CurrentCount = 1
		}
		if CurrentCount == maxValue {
			result = append(result, root.Val)
		} else if CurrentCount > maxValue {
			maxValue = CurrentCount
			result = []int{root.Val}
		}
		travel(root.Right)
	}
	travel(root)
	return result
}

// func findMode(root *TreeNode) []int {
// 	var currentVal, currentCount, maxCount int
// 	var modes []int

// 	inOrderTraversal(root, &currentVal, &currentCount, &maxCount, &modes)

// 	return modes
// }

// func inOrderTraversal(node *TreeNode, currentVal, currentCount, maxCount *int, modes *[]int) {
// 	if node == nil {
// 		return
// 	}

// 	inOrderTraversal(node.Left, currentVal, currentCount, maxCount, modes)

// 	if node.Val == *currentVal {
// 		*currentCount++
// 	} else {
// 		*currentVal = node.Val
// 		*currentCount = 1
// 	}

// 	if *currentCount == *maxCount {
// 		*modes = append(*modes, *currentVal)
// 	} else if *currentCount > *maxCount {
// 		*maxCount = *currentCount
// 		*modes = []int{*currentVal}
// 	}

// 	inOrderTraversal(node.Right, currentVal, currentCount, maxCount, modes)
// }
func main() {

	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 4.0,
	// }
	// var Node5 *TreeNode = &TreeNode{
	// 	Val: 4,
	// }
	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 3,
	// }
	var Node3 *TreeNode = &TreeNode{
		Val: 2,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2, Right: Node3,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1, Right: Node2,
	}
	findMode(Node1)
}
