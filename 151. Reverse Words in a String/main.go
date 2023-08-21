package main

// "fmt"
// "strconv"

/*
Medium
string
done

Given an input string s, reverse the order of the words.
A word is defined as a sequence of non-space characters.
The words in s will be separated by at least one space.
Return a string of the words in reverse order concatenated by a single space.
Note that s may contain leading or trailing spaces or multiple spaces between two words.
The returned string should only have a single space separating the words. Do not include any extra spaces.

Input: s = "the sky is blue"
Output: "blue is sky the"
Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
*/

// func reverseWords(s string) string {
// 	l := 0
// 	p := []byte(s)
// 	for p[0] == ' ' {
// 		l++
// 		p = p[l:]
// 		l = 0
// 	}
// 	r := len(p) - 1
// 	for p[r] == ' ' {
// 		r--
// 		p = p[:r]
// 		r = len(p) - 1

// 	}
// 	for l < r {
// 		p[l], p[r] = p[r], p[l]
// 		l++
// 		r--
// 	}

// 	for i := 0; i < len(p); i++ {
// 		r = i
// 		l = i
// 		for p[r] == ' ' {
// 			p = append(p[0:r], p[r+1:]...)

// 		}
// 		l = r

// 		for p[r+1] != ' ' {
// 			r++
// 			if r+1 >= len(p) {
// 				r = len(p) - 1
// 				break
// 			}
// 		}

// 		i = r + 1
// 		for l < r {
// 			p[l], p[r] = p[r], p[l]
// 			l++
// 			r--
// 		}
// 	}
// 	return string(p)
// }

// func reverseWords(s string) string {

// 	p := strings.Split(s, " ")
// 	ans := ""
// 	for i := len(p) - 1; i >= 0; i-- {
// 		if p[i] != "" {
// 			if ans == "" {
// 				ans = p[i]
// 			} else {
// 				ans = ans + " " + p[i]
// 			}
// 		}
// 	}
// 	return ans
// }

func reverseWords(s string) string {
	slow, fast := len(s), len(s)-1
	out := []rune{}
	for fast >= 0 {
		if s[fast] == ' ' {
			if len(out) > 0 {
				out = append(out, ' ')
			}
			out = append(out, []rune(s[fast+1:slow])...)
			for s[fast] == ' ' {
				fast--
				slow = fast + 1

			}
			continue
			// slow = fast
		}
		// fast--
		if fast == 0 && s[fast] != ' ' {
			if len(out) > 0 {
				out = append(out, ' ')
			}
			out = append(out, []rune(s[fast:slow])...)
		}
		fast--
	}
	return string(out)
}

func reverseWords(s string) string {

	out := []rune{}

	lastIndex := -1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			if lastIndex == -1 {
				lastIndex = i
			}
		} else {
			if lastIndex != -1 {

				if len(out) > 0 {
					out = append(out, ' ')
				}
				out = append(out, []rune(s[i+1:lastIndex+1])...)
				lastIndex = -1
			}
		}
	}

	if lastIndex != -1 {
		if len(out) > 0 {
			out = append(out, ' ')
		}
		out = append(out, []rune(s[0:lastIndex+1])...)
	}

	return string(out)
}
func main() {
	quet := "  hello world  "
	reverseWords(quet)
	quet = "the sky is blue"
	reverseWords(quet)
	quet = "a good   example"
	reverseWords(quet)
}
