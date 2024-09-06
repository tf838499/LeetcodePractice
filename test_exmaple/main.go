package main

import (
	"fmt"
)

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

}

func getIntersectionNode(n) *ListNode {

}
func main() {
	input := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	// input := ListNode{1, &ListNode{2, nil}}
	// input := ListNode{7, &ListNode{7, nil}}
	ans := removeNthFromEnd(&input, 0)
	fmt.Println(ans)
}
