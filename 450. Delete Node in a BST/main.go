package main

import (
	"fmt"
	// "strconv"
)

/*
Given a root node reference of a BST and a key,
delete the node with the given key in the BST.
Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.
*/

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func deleteNode(root *TreeNode, key int) *TreeNode {

// 	if root == nil {
// 		return root
// 	}

// 	if root.Val == key {
// 		if root.Left == nil && root.Right == nil {
// 			return nil
// 		} else {
// 			minNode := findMinNode(root.Right)
// 			if minNode!=nil{
// 				root.Val = minNode.Val
// 			}
// 			// minNodeVal := minNode.Val
// 			root.Right = deleteNode(root.Right, minNode.Val)
// 			// minNode.Left = root.Left

// 			// minNode.Right = root.Left
// 		}
// 	} else if root.Left == nil {
// 		return deleteNode(root.Right, key)
// 	} else if root.Right == nil {
// 		return deleteNode(root.Left, key)
// 	} else if root.Val < key {
// 		root.Right = deleteNode(root.Right, key)
// 	} else if root.Val > key {
// 		root.Left = deleteNode(root.Left, key)
// 	}

//		return root
//	}
//
//	func findMinNode(root *TreeNode) *TreeNode {
//		if root == nil {
//			return root
//		}
//		if root.Left == nil {
//			return root
//		}
//		return findMinNode(root.Left)
//	}
// func deleteNode(root *TreeNode, key int) *TreeNode {
// 	// If the root is nil, the tree is empty, and there's nothing to delete.
// 	if root == nil {
// 		return root
// 	} else if root.Val < key {
// 		// If the key is greater than the current node's value, we need to search in the right subtree.
// 		root.Right = deleteNode(root.Right, key)
// 	} else if root.Val > key {
// 		// If the key is smaller than the current node's value, we need to search in the left subtree.
// 		root.Left = deleteNode(root.Left, key)
// 	} else {
// 		// If we've found the node with the key, there are three cases to consider:
// 		if root.Left == nil && root.Right == nil {
// 			// 1. If the node has no children, it can be safely removed.
// 			return nil
// 		} else if root.Left == nil {
// 			// 2. If the node has only a right child, replace it with its right child.
// 			return root.Right
// 		} else if root.Right == nil {
// 			// 3. If the node has only a left child, replace it with its left child.
// 			return root.Left
// 		} else {
// 			// 4. If the node has both left and right children, find the minimum node in the right subtree,
// 			//    replace the current node's value with that minimum value, and recursively delete the
// 			//    minimum node in the right subtree.
// 			root.Val = findMin(root.Right)
// 			root.Right = deleteNode(root.Right, root.Val)
// 		}
// 	}
// 	return root
// }

// // Function to find the minimum value in a binary tree.
// func findMin(node *TreeNode) int {
// 	minVal := node.Val

// 	// Traverse to the leftmost node to find the minimum value.
// 	for node != nil {
// 		if node.Val < minVal {
// 			minVal = node.Val
// 		}
// 		node = node.Left
// 	}

//		return minVal
//	}
// func deleteNode(root *TreeNode, key int) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	if root.Val == key {
// 		// do delete operation

// 		// case 1: this is a leaf node, directly delete
// 		if root.Left == nil && root.Right == nil {
// 			return nil
// 		}

// 		// case 2: it has only one child, let the one child to replace it
// 		if root.Left == nil && root.Right != nil {
// 			return root.Right
// 		}
// 		if root.Left != nil && root.Right == nil {
// 			return root.Left
// 		}

// 		// case 3: it has both left and right child
// 		if root.Left != nil && root.Right != nil {
// 			// Found the smallest node on the right to replace it
// 			minSubTreeNode := getMin(root.Right)
// 			leftSubTree := root.Left
// 			rightSubTree := deleteNode(root.Right, minSubTreeNode.Val)
// 			minSubTreeNode.Left = leftSubTree
// 			minSubTreeNode.Right = rightSubTree
// 			return minSubTreeNode
// 		}
// 	}

// 	if root.Val > key {
// 		root.Left = deleteNode(root.Left, key)
// 	}

// 	if root.Val < key {
// 		root.Right = deleteNode(root.Right, key)
// 	}

// 	return root
// }

// func getMin(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	if root.Left == nil {
// 		return root
// 	}

//		return getMin(root.Left)
//	}
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Right == nil {
		return root.Left
	}
	if root.Left == nil {
		return root.Right
	}
	minnode := root.Right
	for minnode.Left != nil {
		minnode = minnode.Left
	}
	root.Val = minnode.Val
	root.Right = deleteNode1(root.Right)
	return root
}

func deleteNode1(root *TreeNode) *TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteNode1(root.Left)
	return root
}

// [5,3,6,2,4,null,7]
//
//				5
//			3		6
//	  2       4        7
//
// [6,3,7,2,4]
//
//		6
//	3		7
//
// 2   4
func main() {
	node6 := &TreeNode{Val: 7}
	// node5 := &TreeNode{Val: 4}
	node4 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 6, Right: node6}
	node2 := &TreeNode{Val: 3, Left: node4}
	node1 := &TreeNode{Val: 5, Left: node2, Right: node3}
	// node1 := &TreeNode{Val: 0}
	// node7 := &TreeNode{Val: 80}
	// node6 := &TreeNode{Val: 60}
	// node5 := &TreeNode{Val: 40}
	// // node4 := &TreeNode{Val: 2}
	// node3 := &TreeNode{Val: 70, Left: node6, Right: node7}
	// node2 := &TreeNode{Val: 30, Right: node5}
	// node1 := &TreeNode{Val: 50, Left: node2, Right: node3}
	// [50,30,70,null,40,60,80]
	// p := deleteNode(node1, 50)
	p := deleteNode(node1, 3)
	fmt.Println(p)
}

//			50
//     30        70
//   null  40   60  80
//
//
//
