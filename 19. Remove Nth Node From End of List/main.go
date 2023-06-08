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

// 19
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return nil
// 	}
// 	tmp := head
// 	prev_n := 1
// 	var prev *ListNode
// 	// prev := head
// 	for head != nil {
// 		if head.Next != nil {
// 			if prev_n >= n {
// 				if prev == nil {
// 					prev = tmp
// 				} else {
// 					prev = prev.Next
// 				}
// 			} else {
// 				prev_n++
// 			}
// 			head = head.Next
// 		} else if prev == nil {
// 			tmp = tmp.Next
// 			break
// 		} else {
// 			prev.Next = prev.Next.Next
// 			break
// 		}

// 	}
// 	return tmp
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	dumy := &ListNode{Next: head}
	slow, first := dumy, dumy

	for i := 0; i <= n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dumy
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
