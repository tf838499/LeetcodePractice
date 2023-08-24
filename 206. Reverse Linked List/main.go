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

// func reverseList(head *ListNode) *ListNode {
// 	cur := head
// 	var pre *ListNode
// 	for cur != nil {

// 		tmp := cur.Next
// 		cur.Next = pre
// 		pre = cur
// 		cur = tmp
// 	}
// 	return pre
// }
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var late *ListNode
	for head != nil {
		prev = head.Next
		head.Next = late
		late = head
		head = prev
		// late =
	}

	return late
}

// func reverseListRecursion(pre *ListNode, cur *ListNode) *ListNode {
// 	if cur == nil {
// 		return pre
// 	}
// 	tmp := cur.Next
// 	cur.Next = pre
// 	pre = cur
// 	// reverseListRecursion(pre, tmp)
// 	return reverseListRecursion(pre, tmp)
// }
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
	// outputMessage(HeadNode, "input : ")
	p := reverseList(HeadNode)
	fmt.Println(p)

}
