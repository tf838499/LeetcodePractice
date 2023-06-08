package main

// "fmt"
// "strconv"

/*
Easy
Done

Given the head of a linked list and an integer val,
remove all the nodes of the linked list that has Node.val == val,
and return the new head.

Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]
Input: head = [], val = 1
Output: []
*/

type ListNode struct {
	Val   int
	Index int
	Next  *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	slow, first := &ListNode{}, head

	for first != nil {
		if first.Val == val {
			if slow.Val == 0 {
				head = head.Next
			} else {
				slow.Next = first.Next

			}

		} else {
			slow = first

		}
		first = first.Next
	}
	return head
}
func main() {
	// numtest1_1 := 19

	input := ListNode{7, 1, &ListNode{7, 2, &ListNode{7, 3, &ListNode{7, 4, &ListNode{7, 5, &ListNode{6, 6, nil}}}}}}
	// words := ["bella","label","roller"]
	removeElements(&input, 6)

}
