package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

/*
Medium
linklist
stack
done

You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
*/
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type Myqueue struct {
	Aqueue []ListNode
}

func NewMyqueue() *Myqueue {
	``
	return &Myqueue{
		Aqueue: make([]ListNode, 0),
	}
}
func (m *Myqueue) Push(val *ListNode) {
	m.Aqueue = append(m.Aqueue, *val)
}

// func reorderList(head *ListNode)  {

// }
func reorderList(head *ListNode) {
	// fmt.Println(head)
	Aqueue := NewMyqueue()
	tmp := head.Next
	for tmp != nil {
		Aqueue.Push(tmp)
		tmp = tmp.Next
	}
	for i := len(Aqueue.Aqueue) - 1; i > 0; i-- {
		Aqueue.Aqueue[i].Next = &Aqueue.Aqueue[i-1]
		if i == 1 {
			Aqueue.Aqueue[i].Next.Next = nil
		}
	}
	head.Next = &Aqueue.Aqueue[len(Aqueue.Aqueue)-1]
	fmt.Println(tmp)
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
	reorderList(HeadNode)
}
