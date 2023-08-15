package main

import "fmt"

/*
*
Easy
iterative
linkedlist
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

// func reverseListRecursion(head, last *ListNode) *ListNode {
// 	/* */
// 	if last == nil {
// 		return head
// 	}
// 	temp := last.Next
// 	last.Next = head
// 	return reverseListRecursion(last, temp)
// }
// func reverseList(head *ListNode) *ListNode {
//     if head == nil {
// 		return head
// 	}
// 	var ans *ListNode
// 	temp := head.Next
// 	head.Next = nil
// 	ans = reverseListRecursion(head, temp)
// 	return ans
// }

func reverseList(head *ListNode) *ListNode {
	return reverseListRecursion(head, nil)
}
func reverseListRecursion(head *ListNode, Last *ListNode) *ListNode {
	var tmp *ListNode

	if head.Next != nil {
		if head.Next.Next != nil {
			tmp = reverseListRecursion(head.Next.Next, head.Next)
			head.Next, head.Next.Next = Last, head
		}
		// tmp = reverseListRecursion(head.Next.Next, head.Next)
		// head.Next, head.Next.Next = Last, head
		// fmt.Println(tmp)
	} else if head.Next == nil {
		head.Next = Last
		return head
	}
	return tmp
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
	SecondNode.Next = nil
	thirdNode.Next = Node4
	Node4.Next = Node5
	// outputMessage(HeadNode, "input : ")
	p := reverseList(HeadNode)
	fmt.Println(p)

}
