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

// 142

// func detectCycle(head *ListNode) *ListNode {

//		if head == nil || head.Next == nil {
//			return head
//		}
//		queue := make([]ListNode, 0, 1)
//		pos := -1
//		var ans *ListNode
//		for head != nil && pos == -1 {
//			for i := 0; i < len(queue); i = i + 1 {
//				if *head == queue[i] {
//					pos = i + 1
//					ans = &queue[i-1]
//					break
//				}
//			}
//			queue = append(queue, *head)
//			head = head.Next
//		}
//		if pos == 1 {
//			ans = nil
//		} else {
//			ans.Next = nil
//		}
//		return ans
//	}
// func detectCycle(head *ListNode) *ListNode {
// 	slow, fast := head, head

// 	for fast != nil && fast.Next != nil {
// 		slow, fast = slow.Next, fast.Next.Next
// 		if slow == fast {
// 			slow = head
// 			for slow != fast {
// 				slow, fast = slow.Next, fast.Next
// 			}
// 			return slow
// 		}
// 	}

// 	return nil
// }

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// func detectCycle(head *ListNode) *ListNode {
// 	is_cycle, fast := hasCycle(head)
// 	if is_cycle == true {
// 		for head != fast {
// 			head = head.Next
// 			fast = fast.Next
// 		}
// 		return head
// 	}
// 	return nil
// }

//	func hasCycle(head *ListNode) (bool, *ListNode) {
//		slow := head
//		fast := head
//		for fast != nil && fast.Next != nil {
//			slow = slow.Next
//			fast = fast.Next.Next
//			if slow == fast {
//				return true, slow
//			}
//		}
//		return false, slow
//	}
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
	ans := detectCycle(HeadNode)
	fmt.Println(ans)
	// outputMessage(ans, "result : ")

}

// {2 0xc000116220}
