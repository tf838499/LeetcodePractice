package main

// "fmt"
// "strconv"

/*


 */
type Queue struct {
	myQueue []int
}

func MyQueue() *Queue {
	return &Queue{
		myQueue: []int{},
	}
}
func (q *Queue) Push(x int) {
	for len(q.myQueue) != 0 && x > q.myQueue[len(q.myQueue)-1] {
		// q.myQueue = q.myQueue[1:]
		q.myQueue = q.myQueue[:len(q.myQueue)-1]
	}
	q.myQueue = append(q.myQueue, x)
}
func (q *Queue) Front() int {
	return q.myQueue[0]
}
func (q *Queue) Pop(x int) {
	if len(q.myQueue) != 0 && q.myQueue[0] == x {
		q.myQueue = q.myQueue[1:]
	}
	// else {
	// 	q.myQueue = q.myQueue[0 : len(q.myQueue)-1]
	// }
}
func maxSlidingWindow(nums []int, k int) []int {
	p := MyQueue()
	for i := 0; i < k; i++ {
		p.Push(nums[i])
	}
	tmp := []int{}
	tmp = append(tmp, p.Front())
	for i := k; i < len(nums); i++ {
		p.Pop(nums[i-k])
		p.Push(nums[i])
		tmp = append(tmp, p.Front())
	}
	return tmp
}

// func maxSlidingWindow(nums []int, k int) []int {

// 	result := make([]int, 0)
// 	q := make([]int, 0)
// 	for i := range nums {
// 		if len(q) > 0 && nums[i] > q[0] {
// 			q = q[1:]
// 		}
// 		// result = append(result, q[0])
// 		q = append(q, nums[i])
// 		if q[0] < nums[i] {
// 			q[0] = nums[i]
// 			q = q[0 : len(q)-1]
// 		}
// 		if i >= k-1 {
// 			result = append(result, q[0])
// 		}
// 	}
// 	return result
// }
// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [3,3,5,5,6,7]

//Input: nums = [1,3,1,2,0,5] k = 3
//output: [3,3,2,5]
func main() {
	// s := []int{1, 3, -1, -3, 5, 3, 6, 7}
	// s := []int{1, -1}
	// s := []int{7, 2, 4}
	s := []int{1, 3, 1, 2, 0, 5}
	maxSlidingWindow(s, 3)

}
