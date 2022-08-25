package main

// "fmt"
// "strconv"

/*

Given a string s, find the length of the longest substring without repeating characters.

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

*/

func lengthOfLongestSubstring(s string) int {
	data := []rune(s)
	var slow int = 0
	var fast int = 0
	t := make(map[int32]int)
	var ans = 0
	for fast < len(data) {
		t[data[fast]]++
		for t[data[fast]] > 1 {
			t[data[slow]]--
			slow++
		}
		fast++
		if fast-slow > ans {
			ans = fast - slow
		}

	}

	return ans
}

// func lengthOfLongestSubstring(s string) int {
//     length := len(s)
//     if length <= 1 {
//         return length
//     }
//     longest, p1, m := 0, 0, make(map[byte]int)
//     var char byte
//     for p2 := 0; p2 < len(s); p2++{
//         char = s[p2]
//         if prev, exists := m[char]; exists {
//             if(prev >= p1){
//                 p1 = prev + 1
//             }
//         }
//         m[char] = p2
//         if d := p2 - p1 + 1; d > longest{
//             longest = d
//         }
//     }
//     return longest
// }

// func lengthOfLongestSubstring(s string) int {
// 	i, j, l, r, max := 0, 0, 0, 0, 0

// 	for i = 0; i < len(s); i++ {
// 		repeating := false
// 		for j = l; j < r && repeating == false; j++ {
// 			repeating = s[i] == s[j]
// 		}

// 		if repeating {
// 			l = j
// 		}
// 		r++

// 		if r-l > max {
// 			max = r - l
// 		}
// 	}

// 	return max
// }
func main() {
	/*
		key string has max ASCII code 128
	*/
	// s := "dvdf"
	// s := "abcabcbb"
	// s := "bbbbb"
	// s := "au"
	// s := "pwwkew"
	// s := "tmmzuxt"
	s := "nfpdmpiy"
	// pwwkew
	lengthOfLongestSubstring(s)
	// fmt.Println(l1)

}
