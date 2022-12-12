package main

import (
// "fmt"
// "strconv"
)

/*
Given the head of a linked list and a value x,
partition it such that all nodes less than x come before nodes greater
than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Input: head = [1,4,3,2,5,2], x = 3
Output: [1,2,2,4,3,5]
123452
122345
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// func partition(head *ListNode, x int) *ListNode {

// 	if head == nil {
// 		return head
// 	}
// 	fast, slow := head, head
// 	// ans := head
// 	for slow.Val >= x && slow.Next != nil {
// 		slow = slow.Next
// 		if slow.Val < x {

// 			fast.Next = slow.Next
// 			slow.Next = head
// 			head = slow
// 			break
// 		}
// 		fast = fast.Next
// 	}
// 	ans := head
// 	for slow.Next != nil {
// 		for fast.Next != nil && fast.Next.Val >= x {
// 			fast = fast.Next
// 		}
// 		if fast.Next == nil {
// 			break
// 		}
// 		if slow.Next.Val >= x {
// 			head = slow.Next
// 			slow.Next = fast.Next
// 			fast.Next = fast.Next.Next
// 			// head.Next = fast
// 			slow = slow.Next
// 			slow.Next = head

// 		} else if slow.Next.Val < x {
// 			slow = slow.Next
// 			fast = fast.Next
// 		}

// 	}
// 	// ans = head
// 	return ans
// }
func partition(head *ListNode, x int) *ListNode {
	first, second := &ListNode{}, &ListNode{}
	dummyHeadFirst, dummyHeadSecond := first, second
	//  1,4,3,2,5,2
	for head != nil {
		if head.Val < x {
			first.Next = head
			first = first.Next
		} else {
			second.Next = head
			second = second.Next
		}
		head = head.Next
	}
	first.Next, second.Next = nil, nil // cut the connection
	if dummyHeadSecond.Next != nil {
		first.Next = dummyHeadSecond.Next
	}
	return dummyHeadFirst.Next
}

func main() {
	var Node6 *ListNode = &ListNode{
		Val: 2,
	}
	var Node5 *ListNode = &ListNode{
		Val: 5, Next: Node6,
	}
	var Node4 *ListNode = &ListNode{
		Val: 2, Next: Node5,
	}
	var Node3 *ListNode = &ListNode{
		Val: 3, Next: Node4,
	}
	var Node2 *ListNode = &ListNode{
		Val: 4, Next: Node3,
	}
	var Node1 *ListNode = &ListNode{
		Val: 1, Next: Node2,
	}
	// [2,1]  1,4,3,2,5,2
	partition(Node1, 3)
}
