package main

// "fmt"
// "strconv"
func topKFrequent(nums []int, k int) []int {
	frequementMap := make(map[int]int)
	queue := []int{}
	for _, v := range nums {
		frequementMap[v]++
	}

	for p, v := range frequementMap {
		if len(queue) == 0 {
			queue = append(queue, p)
			continue
		}
		i := 0
		for len(queue) < i && v < frequementMap[queue[i]] {
			i++
		}
		tmp := queue[i:]
		queue = append(queue[:i], p)
		queue = append(queue, tmp...)
	}
	return queue[:k]
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	// numtest2_1 := []int{4, 9, 5}
	// nums = [1,1,1,2,2,3], k = 2
	// words := ["bella","label","roller"]
	topKFrequent(nums, k)

}
