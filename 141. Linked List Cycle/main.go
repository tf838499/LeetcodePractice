package main

import (
	"fmt"
)

/*
*
Medium
two pointer
linkedkist
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

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
func main() {
	var HeadNode *ListNode = &ListNode{
		Val: 3,
	}
	var SecondNode *ListNode = &ListNode{
		Val: 2,
	}
	var thirdNode *ListNode = &ListNode{
		Val: 0,
	}
	var Node4 *ListNode = &ListNode{
		Val: -4,
	}

	// var langs [4]int

	HeadNode.Next = SecondNode
	SecondNode.Next = thirdNode
	thirdNode.Next = Node4
	Node4.Next = SecondNode
	// a := &Node5.Next
	// langs[0] = a
	// b := &Node4.Next
	// fmt.Printf("a : %T\n", a)
	// fmt.Println(b)
	// outputMessage(HeadNode, "input : ")
	ans := hasCycle(HeadNode)
	fmt.Println(ans)
	// outputMessage(ans, "result : ")

}

// {2 0xc000116220}
