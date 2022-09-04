package main

import "fmt"

// "fmt"
// "strconv"

/*
Implement a first in first out (FIFO) queue using only two stacks.
The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:

void push(int x) Pushes element x to the back of the queue.
int pop() Removes the element from the front of the queue and returns it.
int peek() Returns the element at the front of the queue.
boolean empty() Returns true if the queue is empty, false otherwise.

Notes:

You must use only standard operations of a stack, which means only push to top,
peek/pop from top, size, and is empty operations are valid.

Depending on your language, the stack may not be supported natively.
You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type MyQueue struct {
	// Stack *[]int
	Queue *[]int
	front int
	back  int
}

// 複習建構子 queue * & stack
/** Initialize your data structure here. */
func Constructor() MyQueue {
	tmp2 := []int{}
	return MyQueue{Queue: &tmp2, front: 0, back: 0}
}

func (this *MyQueue) Push(x int) {
	*this.Queue = append(*this.Queue, x)
	this.back++
}

func (this *MyQueue) Pop() int {
	if this.back == 0 {
		return this.back
	}
	ans := (*this.Queue)[this.front]
	*this.Queue = (*this.Queue)[this.front+1 : this.back]
	this.back--
	return ans
}

func (this *MyQueue) Peek() int {
	if this.back == 0 {
		return this.back
	}
	return (*this.Queue)[this.front]
}

func (this *MyQueue) Empty() bool {
	return this.back == 0
}

// func (this *MyQueue) Push(x int) {
// 	*this.Stack = append(*this.Stack, x)
// 	fmt.Println(*this.Stack)
// }

// func (this *MyQueue) Pop() int {
// 	if len(*this.Queue) == 0 {
// 		this.fromStackToQueue(this.Stack, this.Queue)
// 	}

// 	popped := (*this.Queue)[len(*this.Queue)-1]
// 	*this.Queue = (*this.Queue)[:len(*this.Queue)-1]
// 	return popped
// }

// func (this *MyQueue) Peek() int {
// 	if len(*this.Queue) == 0 {
// 		this.fromStackToQueue(this.Stack, this.Queue)
// 	}
// 	return (*this.Queue)[len(*this.Queue)-1]
// }

// func (this *MyQueue) Empty() bool {
// 	return len(*this.Stack)+len(*this.Queue) == 0
// }

// func (this *MyQueue) fromStackToQueue(s, q *[]int) {
// 	for len(*s) > 0 {
// 		popped := (*s)[len(*s)-1]
// 		*s = (*s)[:len(*s)-1]

// 		*q = append(*q, popped)
// 	}
// }
func main() {

	// fmt.Println(l1)
	// a:= {"MyQueue", "push", "push", "peek", "pop", "empty"}
	obj := Constructor()
	fmt.Println(obj)
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Pop()
	obj.Push(2)
	// param_2 := obj.Pop();
	// param_3 := obj.Peek();
	// param_4 := obj.Empty();
}
