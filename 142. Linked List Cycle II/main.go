package main

import (
	"fmt"
)

/**
Medium
two pointer
linkedkist
done

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

// 142

// func detectCycle(head *ListNode) *ListNode {

// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	queue := make([]ListNode, 0, 1)
// 	pos := -1
// 	var ans *ListNode
// 	for head != nil && pos == -1 {
// 		for i := 0; i < len(queue); i = i + 1 {
// 			if *head == queue[i] {
// 				pos = i + 1
// 				ans = &queue[i-1]
// 				break
// 			}
// 		}
// 		queue = append(queue, *head)
// 		head = head.Next
// 	}
// 	if pos == 1 {
// 		ans = nil
// 	} else {
// 		ans.Next = nil
// 	}
// 	return ans
// }
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow, fast = slow.Next, fast.Next
			}
			return slow
		}
	}

	return nil
}

// func outputMessage(head *ListNode, title string) {
// 	foo := title
// 	for head != nil {
// 		foo = foo + strconv.Itoa(head.Val) + " "
// 		head = head.Next
// 	}
// 	fmt.Print(foo + "\n")
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
	// var langs [4]int

	HeadNode.Next = SecondNode
	SecondNode.Next = thirdNode
	thirdNode.Next = Node4
	Node4.Next = Node5
	Node5.Next = Node4
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
