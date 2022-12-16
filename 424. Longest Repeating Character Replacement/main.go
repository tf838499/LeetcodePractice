package main

import (
// "fmt"
// "strconv"
// "math"
)

/*
You are given a string s and an integer k.
You can choose any character of the string and
change it to any other uppercase English character.
You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
*/
/*
定义start和end两个标记，中间的内容即是窗口，
计算窗口内所有字母出现的次数，因为全是大写字母，
所以可以用一个26位的数组来记录窗口内每个字母出现的次数。
找到窗口内出现最多的次数，加上允许替换的字母数k，
看是否超过窗口宽度，如果超过了，说明窗口还可以更长，
也就是说窗口内重复的字母的长度可以更长，就将end右移一位，
形成新的窗口，然后继续重复上面的步骤。如果没超过，
说明能构成的最长的重复字母长度已经到顶了，
这时应该将start右移一位，来寻找新的可能的更长重复字母长度。
每次计算重复字母长度时，当出现更长的可能时，都更新最终的结果。
*/
func characterReplacement(s string, k int) int {
	count := make([]int, 128)
	for i := 0; i < len(count); i++ {
		count[i] = 0
	}
	max := 0
	start := 0
	for end := 0; end < len(s); end++ {
		count[s[end]] += 1
		if max < count[s[end]] {
			max = count[s[end]]
		}
		if max+k <= end-start {
			count[s[start]] -= 1
			start += 1
		}
	}
	return len(s) - start
}

func main() {
	a := "AABABBA"
	// a := []int{1, 1, 1, 1, 1, 1, 1, 1}
	characterReplacement(a, 1)
}
