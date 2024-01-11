package main

/*
Given a string s, partition s such that every
substring of the partition is a
palindrome
. Return all possible palindrome partitioning of s.



Example 1:
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Example 2:
Input: s = "a"
Output: [["a"]]


Constraints:

1 <= s.length <= 16
s contains only lowercase English letters.
*/
func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{{}}
	}
	if len(s) == 1 {
		return [][]string{{s}}
	}
	result := [][]string{}
	for i := 0; i < len(s); i++ {
		candidate := s[:i+1]
		if palindrome(candidate) {
			remaining := partition(s[i+1:])
			for _, r := range remaining {
				partition := append([]string{candidate}, r...)
				result = append(result, partition)
			}
		}
	}
	return result
}
func palindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return palindrome(s[1 : len(s)-1])
}

// func partition(s string) [][]string {
// 	if len(s) == 0 {
// 		return [][]string{{}}
// 	}
// 	if len(s) == 1 {
// 		return [][]string{{s}}
// 	}

// 	allPartitions := make([][]string, 0, len(s))
// 	for i := 0; i < len(s); i++ {
// 		candidate := s[:i+1]
// 		if palindrome(candidate) {
// 			remaining := partition(s[i+1:])
// 			for _, r := range remaining {
// 				partition := append([]string{candidate}, r...)
// 				allPartitions = append(allPartitions, partition)
// 			}
// 		}
// 	}

// 	return allPartitions
// }

// func palindrome(s string) bool {
// 	if len(s) <= 1 {
// 		return true
// 	}
// 	if s[0] != s[len(s)-1] {
// 		return false
// 	}
// 	return palindrome(s[1 : len(s)-1])
// }
func main() {
	// isPalindrome("A man, a plan, a canal: Panama")
	partition("abbab")
}
