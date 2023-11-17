package main

// "fmt"
// "strconv"
// "strings"
// "math"

/*
Input: root = [1,2,3,4,null,5,6,null,null,7]
Output: 7

Given the root of a binary tree,
return the leftmost value in the last row of the tree.

Input: root = [1,3,2,5,3,null,9]
Output: [1,3,9]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	ans := 0
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		lenStack := len(stack)
		for i := 0; i < lenStack; i++ {
			nowRoot := stack[i]
			if i == 0 {
				ans = nowRoot.Val
			}
			if nowRoot.Left != nil {
				stack = append(stack, nowRoot.Left)
			}
			if nowRoot.Right != nil {
				stack = append(stack, nowRoot.Right)
			}
		}
		stack = stack[lenStack:]
	}
	return ans
}

// [1,2,3,4,null,5,6,null,null,7]
func main() {
	var Node7 *TreeNode = &TreeNode{
		Val: 7,
	}
	var Node6 *TreeNode = &TreeNode{
		Val: 6,
	}
	var Node5 *TreeNode = &TreeNode{
		Val: 5, Left: Node7,
	}
	var Node4 *TreeNode = &TreeNode{
		Val: 4, Left: Node5, Right: Node6,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 3,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2, Left: Node4,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1, Left: Node2, Right: Node3,
	}

	findBottomLeftValue(Node1)
}
