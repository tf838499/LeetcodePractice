package main

import (
// "fmt"
// "strconv"
// "strings"
// "math"
)

/*
Input: root = [3,9,20,null,null,15,7]
Output: [[15,7],[9,20],[3]]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	// tmp := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for root != nil && len(stack) != 0 {
		stackLen := len(stack)
		tmp := make([]int, 0)
		for i := 0; i < stackLen; i++ {
			top := stack[i]
			if top == nil {
				continue
			}
			tmp = append(tmp, top.Val)
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
			// stack = append(stack, top.Left)
		}
		stack = stack[stackLen:]
		if len(tmp) != 0 {
			ans = append(ans, tmp)
		}
	}
	for l, r := 0, len(ans)-1; l < r; {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}
	return ans
}

// func levelOrderBottom(root *TreeNode) (res [][]int) {
// 	res = [][]int{}
// 	if root == nil {
// 		return
// 	}
// 	q := []*TreeNode{root}
// 	for len(q) > 0 {
// 		level := []int{}
// 		for _, x := range q {
// 			q = q[1:]
// 			if x.Left != nil {
// 				q = append(q, x.Left)
// 			}
// 			if x.Right != nil {
// 				q = append(q, x.Right)
// 			}
// 			level = append(level, x.Val)
// 		}
// 		res = append(res, level)
// 	}
// 	for l, r := 0, len(res)-1; l < r; {
// 		res[l], res[r] = res[r], res[l]
// 		l++
// 		r--
// 	}
// 	return
// }
func main() {
	var Node5 *TreeNode = &TreeNode{
		Val: 7.0,
	}
	var Node4 *TreeNode = &TreeNode{
		Val: 15.0,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 20.0, Right: Node5, Left: Node4,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 9.0,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 3.0, Right: Node3, Left: Node2,
	}

	levelOrderBottom(Node1)
}
