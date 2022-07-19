package main

import (
	"fmt"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 206
func reverseList(head *ListNode) *ListNode {
	/*
		ForwardNode : record pre node
		CurrentNode : record next node
		method		: [first head equal current ，then current node pass next node ，
		then head next node directed forward node (reverse)，complete once reverse，go back first and continus reverse]
	*/
	var ForwardNode *ListNode
	var CurrentNode *ListNode
	CurrentNode = head

	for CurrentNode != nil {
		head = CurrentNode
		CurrentNode = head.Next
		head.Next = ForwardNode
		ForwardNode = head
	}
	return head
}

// func reverseListRecursion(head *ListNode) *ListNode {
// 	/*
// 		pass
// 	*/
// 	return nil
// }
func outputMessage(head *ListNode, title string) {
	foo := title
	for head != nil {
		foo = foo + strconv.Itoa(head.Val) + " "
		head = head.Next
	}
	fmt.Print(foo + "\n")
}
func main() {
	var HeadNode *ListNode = &ListNode{
		Val: 1.0,
	}
	var SecondNode *ListNode = &ListNode{
		Val: 2.0,
	}
	var thirdNode *ListNode = &ListNode{
		Val: 3.0,
	}
	var Node4 *ListNode = &ListNode{
		Val: 4.0,
	}
	var Node5 *ListNode = &ListNode{
		Val: 5.0,
	}

	HeadNode.Next = SecondNode
	SecondNode.Next = thirdNode
	thirdNode.Next = Node4
	Node4.Next = Node5
	outputMessage(HeadNode, "input : ")
	ans := reverseList(HeadNode)
	outputMessage(ans, "result : ")

}
