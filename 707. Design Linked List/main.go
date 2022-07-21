package main

import (
	"fmt"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MyLinkedList struct {
	head *ListNode1
	// head ListNode // 會變初始化建立一個val = 0 Next = nil

}
type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func Constructor() MyLinkedList {
	objLinkedList := new(MyLinkedList)
	return *objLinkedList
}

func (this *MyLinkedList) Get(index int) int {
	tmp := this.head
	counter := 0
	for tmp != nil {
		if index == counter {
			return tmp.Val
		}
		tmp = tmp.Next
		counter++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	addHead := ListNode1{Val: val}
	addHead.Next = this.head
	this.head = &addHead
}

func (this *MyLinkedList) AddAtTail(val int) {
	tmp := this.head
	if this.head == nil {
		this.AddAtHead(val)
	} else {
		for this.head.Next != nil {
			this.head = this.head.Next
		}
		this.head.Next = &ListNode1{Val: val}
		this.head = tmp
	}

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

	tmp := this.head
	counter := 0

	if index == 0 && tmp == nil {
		this.head = &ListNode1{Val: val, Next: nil}
		return
	} else if index == 0 {
		this.head = &ListNode1{Val: val, Next: tmp}
		return
	}
	for index != counter {
		counter++
		if index == counter {
			if tmp == nil {
				return
			} else if tmp.Next != nil {
				tmp.Next = &ListNode1{Val: val, Next: tmp.Next}
			} else {
				tmp.Next = &ListNode1{Val: val}
			}
			return
		}
		tmp = tmp.Next
	}
	return
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	tmp := this.head
	counter := 0
	var prev *ListNode1
	for tmp != nil {
		if index == counter {
			if prev == nil {
				this.head = tmp.Next
				break
			}
			if prev.Next.Next == nil {
				prev.Next = nil
			} else {
				prev.Next = prev.Next.Next
			}
			break
		}
		prev = tmp
		tmp = tmp.Next
		counter++
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func outputMessage(head *ListNode1, title string) {
	foo := title
	for head != nil {
		foo = foo + strconv.Itoa(head.Val) + " "
		head = head.Next
	}
	fmt.Print(foo + "\n")
}
func main() {

	obj := Constructor()
	// obj.AddAtHead(1)
	// obj.AddAtTail(3)
	obj.AddAtIndex(0, 1)
	obj.AddAtIndex(1, 2)
	// obj.AddAtIndex(1, 3)
	// obj.Get(1)
	// obj.DeleteAtIndex(0)
	obj.Get(0)

	// outputMessage(obj, "input : ")

}
