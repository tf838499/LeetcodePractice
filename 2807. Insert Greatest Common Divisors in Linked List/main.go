package main

import (
	"fmt"
	// "strconv"
)

/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list.
The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	cur := head
	for cur.Next != nil {
		a := cur.Val
		b := cur.Next.Val
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		tmp := cur.Next
		cur.Next = &ListNode{Val: a, Next: tmp}
		cur = tmp

	}
	return head
}

func main() {

	var Node1 *ListNode = &ListNode{
		Val: 18, Next: &ListNode{
			Val: 6, Next: &ListNode{
				Val: 10, Next: &ListNode{
					Val: 3,
				},
			},
		},
	}

	p := insertGreatestCommonDivisors(Node1)
	fmt.Println(p)
}
