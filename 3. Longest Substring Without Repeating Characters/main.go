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
