package main

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

// 206
// func reverseList(head *ListNode) *ListNode {
// 	/*
// 		ForwardNode : record pre node
// 		CurrentNode : record next node
// 		method		: [first head equal current ，then current node pass next node ，
// 		then head next node directed forward node (reverse)，complete once reverse，go back first and continus reverse]
// 	*/
// 	var ForwardNode *ListNode
// 	var CurrentNode *ListNode
// 	CurrentNode = head

// 	for CurrentNode != nil {
// 		head = CurrentNode
// 		CurrentNode = head.Next
// 		head.Next = ForwardNode
// 		ForwardNode = head
// 	}
// 	return head
// }
// func reverseList(head *ListNode) *ListNode {
//     if head == nil{
// 		return nil
// 	}
// 	temp := head.Next
// 	head.Next = nil
// 	return reversePair(head, temp)
// }

//	func reversePair(last *ListNode, head *ListNode) *ListNode {
//		if head == nil{
//			return last
//		}
//		temp := head.Next
//		head.Next = last
//		return reversePair(head, temp)
//	}
func reverseListRecursion(head, last *ListNode) *ListNode {
	/* */
	if last == nil {
		return head
	}
	temp := last.Next
	last.Next = head
	return reverseListRecursion(last, temp)
}

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	var ans *ListNode
// 	temp := head.Next
// 	head.Next = nil
// 	ans = reverseListRecursion(head, temp)
// 	return ans
// }

//	func reverseList(head *ListNode) *ListNode {
//		help(head, nil)
//		return head
//	}
//
func help(head *ListNode, pre *ListNode) *ListNode {
	if head != nil {
		head = head.Next
		rever := help(head, head.Next)
		pre = rever
	}
	return head
}
func helper(current *ListNode, prev *ListNode) *ListNode {
	if current == nil {
		return prev
	}
	next := current.Next
	current.Next = prev
	return helper(next, current)
}

func reverseList(head *ListNode) *ListNode {
	return helper(head, nil)
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
	// outputMessage(HeadNode, "input : ")
	reverseList(HeadNode)

}
