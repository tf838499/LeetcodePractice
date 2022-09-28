package main

import (
// "fmt"
// "strconv"
// "strings"
// "math"
)

/*
Medium
bfs
queue
done

Given the root of a binary tree,
imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.

Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
*/

type Node struct {
	Val      int
	Children []*Node
}

// func rightSideView(root *TreeNode) (ans []int) {
// 	stack := []*TreeNode{root}
// 	if root == nil {
// 		return
// 	}
// 	// ans = append(ans, root.Val)
// 	for len(stack) != 0 {
// 		// ans = append(ans, root.Val)
// 		tmp := []int{}
// 		for _, i := range stack {
// 			stack = stack[1:]
// 			if i.Left != nil {
// 				stack = append(stack, i.Left)
// 			}
// 			if i.Right != nil {
// 				stack = append(stack, i.Right)
// 			}
// 			tmp = append(tmp, i.Val)
// 		}
// 		ans = append(ans, tmp[len(tmp)-1])
// 	}
// 	return
// }

func levelOrder(root *Node) (ans [][]int) {
	stack := []*Node{root}
	if root == nil {
		return
	}
	for len(stack) != 0 {
		// ans = append(ans, root.Val)
		tmp := []int{}
		for _, i := range stack {
			stack = stack[1:]
			for j := 0; j < len(i.Children); j++ {
				stack = append(stack, i.Children[j])
			}
			// fmt.Println(i)
			// for j := 0; j < len(i); j++ {
			// 	fmt.Println(i[j])
			// }
			tmp = append(tmp, i.Val)
		}
		ans = append(ans, tmp)
	}

	return
}

func main() {
	var Node6 *Node = &Node{
		Val: 6.0,
	}
	var Node5 *Node = &Node{
		Val: 5.0,
	}
	var Node4 *Node = &Node{
		Val: 4.0,
	}
	var Node3 *Node = &Node{
		Val: 3.0, Children: []*Node{Node5, Node6},
	}
	var Node2 *Node = &Node{
		Val: 2.0,
	}
	var Node1 *Node = &Node{
		Val: 1.0, Children: []*Node{Node2, Node3, Node4},
	}
	// Node1 := Node{Val: 1.0, Children: []*Node{Node2, Node3}}

	levelOrder(Node1)
}
