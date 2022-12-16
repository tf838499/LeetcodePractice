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

// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
// Input: list1 = [], list2 = [0]
// Output: [0]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list2, list1.Next)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

}

// func mergeTwoListsRe(list1 *ListNode, list2 *ListNode, head *ListNode) *ListNode {
// 	var tmp *ListNode
// 	// head = list1
// 	if list1 == nil {
// 		return head
// 	} else if list2 == nil {
// 		return head
// 	}

// 	if list1.Val > list2.Val {
// 		tmp = list2.Next
// 		list2.Next = list1
// 		return mergeTwoListsRe(list2.Next, tmp, head)
// 	} else {
// 		tmp = list1.Next
// 		list1.Next = list2
// 		return mergeTwoListsRe(tmp, list1.Next, head)
// 	}
// }
func main() {
	// var Node3_1 *ListNode = &ListNode{
	// 	Val: 4,
	// }
	// var Node2_1 *ListNode = &ListNode{
	// 	Val: 2, Next: Node3_1,
	// }
	var Node1_1 *ListNode = &ListNode{
		Val: 2,
	}

	// var Node3_2 *ListNode = &ListNode{
	// 	Val: 4,
	// }
	// var Node2_2 *ListNode = &ListNode{
	// 	Val: 3, Next: Node3_2,
	// }
	var Node1_2 *ListNode
	p := mergeTwoLists(Node1_1, Node1_2)
	fmt.Println(p)
}
