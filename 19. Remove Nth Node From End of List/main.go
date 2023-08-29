package main

import "fmt"

/*
*
Medium
linkedlist
two pointer
done
  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	slow, fast := head, head
// 	for i := 0; i < n; i++ {
// 		fast = fast.Next
// 	}

// 	for fast != nil {
// 		fast = fast.Next
// 		slow = slow.Next
// 	}
// 	slow.Next = slow.Next.Next
// 	return head
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummyHead := &ListNode{-1, head}

	cur, prevOfRemoval := dummyHead, dummyHead

	for cur.Next != nil {

		// n step delay for prevOfRemoval
		if n <= 0 {
			prevOfRemoval = prevOfRemoval.Next
		}

		cur = cur.Next

		n -= 1
	}

	// Remove the N-th node from end of list
	nthNode := prevOfRemoval.Next
	prevOfRemoval.Next = nthNode.Next

	return dummyHead.Next

}
func main() {

	var HeadNode *ListNode = &ListNode{
		Val: 1.0,
	}
	// var SecondNode *ListNode = &ListNode{
	// 	Val: 2.0,
	// }
	// var thirdNode *ListNode = &ListNode{
	// 	Val: 3.0,
	// }
	// var Node4 *ListNode = &ListNode{
	// 	Val: 4.0,
	// }
	// var Node5 *ListNode = &ListNode{
	// 	Val: 5.0,
	// }

	// HeadNode.Next = SecondNode
	// SecondNode.Next = thirdNode
	// thirdNode.Next = Node4
	// Node4.Next = Node5

	P := removeNthFromEnd(HeadNode, 1)
	fmt.Println(P)
}
