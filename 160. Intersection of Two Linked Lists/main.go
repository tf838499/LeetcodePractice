package main

// "strconv"

/*

 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//	func getIntersectionNode(headA, headB *ListNode) *ListNode {
//		lenB, lenA := 0, 0
//		curA := headA
//		curB := headB
//		for curA != nil || curB != nil {
//			if curA != nil {
//				lenA++
//				curA = curA.Next
//			}
//			if curB != nil {
//				lenB++
//				curB = curB.Next
//			}
//		}
//		var step int
//		if lenA < lenB {
//			curA, curB = headB, headA
//			step = lenB - lenA
//		} else {
//			curA, curB = headA, headB
//			step = lenA - lenB
//		}
//		for i := 0; i < step; i++ {
//			curA = curA.Next
//		}
//		for curA != curB {
//			curA = curA.Next
//			curB = curB.Next
//		}
//		return curA
//	}
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	curA := headA
	curB := headB
	for curA != curB {
		curA = curA.Next
		curB = curB.Next
		if curA == nil {
			curA = headB
		}
		if curB == nil {
			curB = headA
		}
	}
	return curA
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}
	return l1
}
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
	B := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1}}}
	getIntersectionNode(A, B)
}
