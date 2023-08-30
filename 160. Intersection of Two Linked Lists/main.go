package main

// "strconv"

/*
Given the heads of two singly linked-lists headA and headB,
return the node at which the two lists intersect.
If the two linked lists have no intersection at all, return null.

For example, the following two linked lists begin to intersect at node c1:
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	curA := headA
	curB := headB
	for curA != curB {

		curA = curA.Next
		curB = curB.Next
		if curA == nil && curB == nil {
			return nil
		}
		if curA == nil {
			curA = headB
		}
		if curB == nil {
			curB = headA
		}
	}
	return curA
}

// func getIntersectionNode(headA, headB *ListNode) *ListNode {

// 	curA := headA
// 	curB := headB
// 	for curA != curB {
// 		curA = curA.Next
// 		curB = curB.Next
// 		if curA == nil {
// 			curA = headB
// 		}
// 		if curB == nil {
// 			curB = headA
// 		}
// 	}
// 	return curA
// }
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	l1, l2 := headA, headB
// 	for l1 != l2 {
// 		if l1 != nil {
// 			l1 = l1.Next
// 		} else {
// 			l1 = headB
// 		}

// 		if l2 != nil {
// 			l2 = l2.Next
// 		} else {
// 			l2 = headA
// 		}
// 	}

// 	return l1
// }
func main() {

	var HeadNode *ListNode = &ListNode{
		Val: 8,
	}
	var SecondNode *ListNode = &ListNode{
		Val: 4,
	}
	var thirdNode *ListNode = &ListNode{
		Val: 5,
	}
	HeadNode.Next = SecondNode
	SecondNode.Next = thirdNode
	A := &ListNode{Val: 4, Next: &ListNode{Val: 1}}
	B := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3}}}
	getIntersectionNode(A, B)
}
