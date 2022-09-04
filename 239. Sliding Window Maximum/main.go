package main

// "fmt"
// "strconv"

/*
You are given an array of integers nums,
there is a sliding window of size k which is moving from the very left of the array to the very right.
You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.


Notes:

You must use only standard operations of a stack, which means only push to top,
peek/pop from top, size, and is empty operations are valid.

Depending on your language, the stack may not be supported natively.
You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.
*/

/*
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
*/
// type Node struct {
// 	val  int
// 	Next *Node
// }

// type MyQueueArray struct {
// 	// Stack *[]int
// 	QueueHead []int
// 	front     int
// 	back      int
// 	len       int
// }

// func Constructor() MyQueueArray {
// 	MyQueue := new(MyQueueArray)
// 	return *MyQueue
// }
// func (this *MyQueueArray) Front_value() int {
// 	return this.QueueHead[0]
// }
// func (this *MyQueueArray) Back_value() int {
// 	return this.QueueHead[this.back]
// }

// func (this *MyQueueArray) Front_Push(x int) {
// 	HeadAdd := []int{x}
// 	HeadAdd = append(HeadAdd, this.QueueHead...)
// 	this.back++
// 	this.QueueHead = HeadAdd
// }

// func (this *MyQueueArray) Back_Push(x int) {
// 	this.QueueHead = append(this.QueueHead, x)
// 	this.back++
// }
// func (this *MyQueueArray) Front_Pop() int {
// 	if this.back == 0 {
// 		return 0
// 	}
// 	this.QueueHead = this.QueueHead[1:len(this.QueueHead)]
// 	this.back--
// 	return this.back
// }
// func (this *MyQueueArray) Back_Pop() int {
// 	if this.back == 0 {
// 		return 0
// 	}
// 	this.QueueHead = this.QueueHead[:len(this.QueueHead)-1]
// 	this.back--
// 	return this.back
// }

// func maxSlidingWindow(nums []int, k int) []int {
// 	var ans = []int{}
// 	deque := Constructor()
// 	// stack := make(map[int]int)
// 	for i := 0; i < k; i++ {
// 		if deque.back == 0 {
// 			deque.Front_Push(nums[i])
// 		} else if deque.back < k {
// 			if deque.Front_value() < nums[i] {
// 				deque.Front_Push(nums[i])
// 			} else {
// 				deque.Back_Push(nums[i])

// 			}
// 		}
// 	}
// 	ans = append(ans, deque.Front_value())
// 	for i := k; i <= len(nums); i++ {
// 		// ans = append(ans, deque.Front_value())
// 		if deque.Front_value() < nums[i] {
// 			deque.Front_Push(nums[i])
// 		} else {
// 			deque.Back_Push(nums[i])
// 		}
// 		if nums[i-k] == deque.Front_value() {
// 			deque.Front_Pop()
// 		}

// 		ans = append(ans, deque.Front_value())
// 	}
// 	// for index, value := range nums {
// 	// 	if deque.back == 0 {
// 	// 		deque.Front_Push(value)
// 	// 		continue
// 	// 	} else if deque.back < k {
// 	// 		if deque.Front_value() < value {
// 	// 			deque.Front_Push(value)
// 	// 			continue
// 	// 		} else {
// 	// 			deque.Back_Push(value)
// 	// 			continue
// 	// 		}
// 	// 	}
// 	// 	ans = append(ans, deque.Front_value())
// 	// stack[index] = value

// 	// fmt.Println(index)
// 	// fmt.Println(index)
// 	// fmt.Print(ans[len(ans)-1])
// 	// }
// 	return ans
// }
// 封装单调队列的方式解题
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MyQueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	length := len(nums)
	res := make([]int, 0)
	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	// 记录前k个元素的最大值
	res = append(res, queue.Front())

	for i := k; i < length; i++ {
		// 滑动窗口移除最前面的元素
		queue.Pop(nums[i-k])
		// 滑动窗口添加最后面的元素
		queue.Push(nums[i])
		// 记录最大值
		res = append(res, queue.Front())
	}
	return res
}
func main() {

	// fmt.Println(l1)
	// a:= {"MyQueue", "push", "push", "peek", "pop", "empty"}
	var nums = []int{1, 3, -1, -1, 2, 1, 5, 3, 6, 7}
	// var nums = []int{7, 2, 4}
	k := 4
	maxSlidingWindow(nums, k)
	// param_2 := obj.Pop();
	// param_3 := obj.Peek();
	// param_4 := obj.Empty();
}
