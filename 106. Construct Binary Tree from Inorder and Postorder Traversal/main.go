package main

// "fmt"
// "strconv"
// "strings"
// "math"

/*
Given two integer arrays inorder and postorder where inorder
is the inorder traversal of a binary tree and postorder
is the postorder traversal of the same tree,
construct and return the binary tree.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	med := postorder[len(postorder)-1]
	inIndex := 0
	for inorder[inIndex] != med {
		inIndex++
	}
	postorder = postorder[0 : len(postorder)-1]
	leftNode := buildTree(inorder[0:inIndex], postorder[0:inIndex])
	rightNode := buildTree(inorder[inIndex+1:], postorder[inIndex:])
	return &TreeNode{Val: med, Left: leftNode, Right: rightNode}
}
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}

	pivotId := 0
	for pivotId < n && inorder[pivotId] != postorder[n-1] {
		pivotId++
	}

	root := new(TreeNode)
	root.Val = postorder[n-1]
	root.Left = buildTree(inorder[:pivotId], postorder[:pivotId])
	root.Right = buildTree(inorder[pivotId+1:], postorder[pivotId:n-1])
	return root
}
func main() {
	// inorder := []int{9, 3, 15, 20, 7}
	// postorder := []int{9, 15, 7, 20, 3}
	inorder := []int{2}
	postorder := []int{2}
	buildTree(inorder, postorder)
}
