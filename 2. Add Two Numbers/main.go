package main

import (
// "fmt"
// "strconv"
)

/*
Medium
linked list
done

You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Input: l1 = [0], l2 = [0]
Output: [0]
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans *ListNode
	var tmpNode *ListNode
	var carry bool
	// both := true
	for l1 != nil || l2 != nil {
		tmp := 0
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		if carry == true {
			tmp += 1
		}

		Node := &ListNode{tmp % 10, nil}
		if tmpNode == nil {
			ans = Node
			tmpNode = ans
			// ans = Node
		} else {
			ans.Next = Node
			ans = ans.Next
		}

		carry = tmp/10 >= 1
	}
	if carry == true {
		Node := &ListNode{1, nil}
		ans.Next = Node
	}

	return tmpNode
}
func main() {
	l1_t := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2_t := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	// words := ["bella","label","roller"]
	// commonChars(words)
	// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	addTwoNumbers(l1_t, l2_t)
	// fmt.Println(l1)

}
