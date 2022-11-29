package main

import (
	// "fmt"
	// "strconv"
	"sort"
)

/*


You are given an integer array matches where matches[i] = [winneri, loseri]
indicates that the player winneri defeated player loseri in a match.

Return a list answer of size 2 where:

answer[0] is a list of all players that have not lost any matches.
answer[1] is a list of all players that have lost exactly one match.
The values in the two lists should be returned in increasing order.

Note:

You should only consider the players that have played at least one match.
The testcases will be generated such that no two matches will have the same outcome.

Input: matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
Output: [[1,2,10],[4,5,7,8]]
Explanation:
Players 1, 2, and 10 have not lost any matches.
Players 4, 5, 7, and 8 each have lost one match.
Players 3, 6, and 9 each have lost two matches.
Thus, answer[0] = [1,2,10] and answer[1] = [4,5,7,8].
*/

// func findWinners(matches [][]int) [][]int {
// 	winneri := make(map[int]int)
// 	losserb := make(map[int]int)
// 	for _, i := range matches {
// 		// if winneri[i[0]] == 0 && losserb[i[0]] != true {
// 		winneri[i[0]]++
// 		winneri[i[1]]--
// 		losserb[i[1]]--
// 	}
// 	// sort.Ints(winneri)
// 	winerArray := []int{}
// 	loserArray := []int{}
// 	for i, j := range winneri {
// 		if j >= 1 && losserb[i] == 0 {
// 			winerArray = append(winerArray, i)
// 		} else if losserb[i] < 0 && losserb[i] >= -1 {
// 			loserArray = append(loserArray, i)
// 		}
// 		// loserArray = append(loserArray, i)

// 	}
// 	sort.Ints(winerArray)
// 	sort.Ints(loserArray)
// 	return [][]int{winerArray, loserArray}
// }

// func findWinners(matches [][]int) [][]int {
//     winners, losers := make(map[int]bool), make(map[int]int)
//     for _, match := range matches {
//         winners[match[0]] = true
//         losers[match[1]]++
//     }

//     oneLosses := []int{}
//     for k, v := range losers {
//         if v == 1 {
//             oneLosses = append(oneLosses, k)
//         }
//         if _, ok := winners[k]; ok {
//             delete(winners, k)
//         }
//     }
//     sort.Ints(oneLosses)

//     noLose := []int{}
//     for k := range winners {
//         noLose = append(noLose, k)
//     }
//     sort.Ints(noLose)

//     return [][]int{noLose, oneLosses}
// }
/*
answer[0] - not lost any matches. - lost lost = 0
answer[1] - lost exactly one match. - lost lost = 1
*/
func findWinners(matches [][]int) [][]int {
	players := make(map[int]int)
	// Record player with 0 if won any matches
	// Add 1 for each lost match by player
	for _, match := range matches {
		players[match[0]] += 0
		players[match[1]] += 1
	}

	answer := make([][]int, 2)
	answer[0] = make([]int, 0)
	answer[1] = make([]int, 0)
	for k, v := range players {
		if v == 1 {
			answer[1] = append(answer[1], k)
		} else if v == 0 {
			answer[0] = append(answer[0], k)
		}
	}
	sort.Ints(answer[0])
	sort.Ints(answer[1])
	return answer
}
func main() {
	// [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
	foo := [][]int{[]int{1, 3}, []int{2, 3}, []int{3, 6}, []int{5, 6}, []int{5, 7}, []int{4, 5}, []int{4, 8}, []int{4, 9}, []int{10, 4}, []int{10, 9}}
	// [[1,2],[2,1],[1,3],[3,1],[2,3],[3,2]]
	// foo := [][]int{[]int{1, 2}, []int{2, 1}, []int{1, 3}, []int{3, 1}, []int{2, 3}, []int{3, 2}}
	findWinners(foo)
}
