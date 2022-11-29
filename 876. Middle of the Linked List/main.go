package main

import (
// "fmt"
// "strconv"
)

/*



Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.

nput: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.
1234567
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
*/
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

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	midNode := head
	count := 0
	for head != nil {
		if count%2 == 1 {
			midNode = midNode.Next
		}
		count++
		head = head.Next
	}
	return midNode
}

func main() {
	// var Node6 *ListNode = &ListNode{
	// 	Val: 6.0,
	// }
	var Node5 *ListNode = &ListNode{
		Val: 5.0,
	}
	var Node4 *ListNode = &ListNode{
		Val: 4.0, Next: Node5,
	}
	var Node3 *ListNode = &ListNode{
		Val: 3.0, Next: Node4,
	}
	var Node2 *ListNode = &ListNode{
		Val: 2.0, Next: Node3,
	}
	var Node1 *ListNode = &ListNode{
		Val: 1.0, Next: Node2,
	}
	middleNode(Node1)
}
