package main

import "fmt"

// "fmt"
// "strconv"
// "strings"
// "math"

/*
easy
recursion
dfs
done

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
/* recursive */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	var tmp TreeNode

	res := 0
	for len(queue) != 0 {
		tmp = *queue[len(queue)-1]
		queueSize := len(queue)

		for i := 0; i < queueSize; i++ {
			tmp = *queue[0]

			if i == 0 {
				res++
			}
			if tmp.Left == nil && tmp.Right == nil {
				return res
			}
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
			queue = queue[1:]
		}
	}
	return res
}

// func minDepth(root *TreeNode) int {
//     if root == nil {return 0}

//     var queue []*TreeNode
//     enqueue(&queue, root)

//     var res int

//     for len(queue) > 0 {
//         queueSize := len(queue)

//         for i := 0; i < queueSize; i++ {
//             dequeuedItem := dequeue(&queue)

//             if i == 0 {
//                 res++
//             }

//             if dequeuedItem.Left == nil && dequeuedItem.Right == nil {
//                 return res
//             }

//             if dequeuedItem.Left != nil {
//                 enqueue(&queue, dequeuedItem.Left)
//             }

//             if dequeuedItem.Right != nil {
//                 enqueue(&queue, dequeuedItem.Right)
//             }
//         }
//     }

//     return res
// }

func enqueue(queue *[]*TreeNode, item *TreeNode) {
	if queue == nil {
		panic("nil pointer")
	}

	*queue = append(*queue, item)
}

func dequeue(queue *[]*TreeNode) *TreeNode {
	if queue == nil {
		panic("nil pointer")
	}

	if len(*queue) == 0 {
		panic("empty queue")
	}

	dequeuedItem := (*queue)[0]
	*queue = (*queue)[1:]

	return dequeuedItem
}
func main() {
	// var Node7 *TreeNode = &TreeNode{
	// 	Val: 3.0,
	// }
	// var Node6 *TreeNode = &TreeNode{
	// 	Val: 4.0,
	// }
	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 4.0,
	// }
	var Node4 *TreeNode = &TreeNode{
		Val: 3.0,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 2.0, Left: Node4,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2.0, Left: Node3,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2,
	}

	P := minDepth(Node1)
	fmt.Println(P)
}
