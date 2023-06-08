package main

import "fmt"

/*
*
Medium
linklists
recursion
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

// 24
// func swapPairs(head *ListNode) *ListNode {

// 	CurrentNode := swapPairsIterate(head)

// 	return CurrentNode
// }

// func swapPairsIterate(head *ListNode) *ListNode {
// 	var NodeNext *ListNode

// 	if head != nil && head.Next != nil {
// 		NodeNext = head.Next
// 		head.Next = NodeNext.Next
// 		NodeNext.Next = head
// 		head.Next = swapPairsIterate(head.Next)
// 		return NodeNext
// 	} else {
// 		return head
// 	}

// }
// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	tmp := head.Next
// 	var p *ListNode
// 	if head.Next != nil {
// 		head.Next = tmp.Next
// 		tmp.Next = head
// 		p = swapPairs(head.Next)
// 		tmp.Next.Next = p
// 		return tmp

//		}
//		return head
//	}
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next != nil {
		p := swapPairs(head.Next.Next)
		tmp := head.Next
		head.Next = p
		tmp.Next = head

		return tmp
	}
	return head
}

//	func swapPairsNormal(head *ListNode) *ListNode {
//		/*
//			pass
//		*/
//		return nil
//	}

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

	ans := swapPairs(HeadNode)
	fmt.Println(ans)

}
