package main

import (
	"fmt"
	"strconv"
)

/**
Medium
linkedlist
two pointer
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

// 19
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	tmp := head
	prev_n := 1
	var prev *ListNode
	// prev := head
	for head != nil {
		if head.Next != nil {
			if prev_n >= n {
				if prev == nil {
					prev = tmp
				} else {
					prev = prev.Next
				}
			} else {
				prev_n++
			}
			head = head.Next
		} else if prev == nil {
			tmp = tmp.Next
			break
		} else {
			prev.Next = prev.Next.Next
			break
		}

	}
	return tmp
}

func outputMessage(head *ListNode, title string) {
	foo := title
	for head != nil {
		foo = foo + strconv.Itoa(head.Val) + " "
		head = head.Next
	}
	fmt.Print(foo + "\n")
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
	outputMessage(HeadNode, "input : ")
	ans := removeNthFromEnd(HeadNode, 1)
	outputMessage(ans, "result : ")

}
