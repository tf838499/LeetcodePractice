package main

import "container/heap"

// import "container/heap"

// "fmt"
// "strconv"
// func topKFrequent(nums []int, k int) []int {

// 	// (1) count repetitions ans save them into map
// 	//where key is the number, and value is the count of repetitions
// 	count := make(map[int]int)
// 	for _, num := range nums {
// 		count[num]++
// 	}

// 	// (2) convert map of counts into slice
// 	//where index is the count, and value is the slice of numbers with the same count of repetitions
// 	freqs := make([][]int, 1+len(nums))
// 	for key, val := range count {
// 		freqs[val] = append(freqs[val], key)
// 	}

// 	// (3) go backwards through the slice of frequences and add non-null values to result
// 	// we are going backwards as the index equals frequency, and we need top K most frequent numbers
// 	res := make([]int, 0, k)
// 	for i := len(freqs) - 1; i >= 0; i-- {
// 		if freqs[i] != nil {
// 			res = append(res, freqs[i]...)
// 			if len(res) >= k {
// 				res = res[:k]
// 				break
// 			}
// 		}
// 	}

// 	return res
// }
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func topKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}
	//记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, value := range map_num {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}
func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 5, 5, 5, 5}
	k := 2
	topKFrequent(nums, k)

}
