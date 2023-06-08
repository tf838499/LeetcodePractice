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
// 	sa := strings.Split(s, " ")
// 	var ans string
// 	flag := 0
// 	stringsLen := len(sa)
// 	for i := 0; i < stringsLen; i++ {

//			if sa[stringsLen-(i+1)] != "" {
//				if flag == 1 {
//					ans = ans + " "
//				}
//				ans = ans + sa[stringsLen-(i+1)]
//				flag = 1
//			}
//			// if flag == 1 {
//			// 	ans = ans + sa[stringsLen-(i+1)] + " "
//			// }
//			// if sa[stringsLen-(i+1)] != "" && i != stringsLen-1 {
//			// 	ans = ans + sa[stringsLen-(i+1)] + " "
//			// } else if i == stringsLen-1 && sa[stringsLen-(i+1)] != " " {
//			// 	ans = ans + sa[stringsLen-(i+1)]
//			// }
//		}
//		return ans
//	}
func reverseWords(s string) string {
	l := 0
	p := []byte(s)
	for p[0] == ' ' {
		l++
		p = p[l:]
		l = 0
	}
	r := len(p) - 1
	for p[r] == ' ' {
		r--
		p = p[:r]
		r = len(p) - 1

	}
	for l < r {
		p[l], p[r] = p[r], p[l]
		l++
		r--
	}

	for i := 0; i < len(p); i++ {
		r = i
		l = i
		for p[r] == ' ' {
			p = append(p[0:r], p[r+1:]...)

		}
		l = r

		for p[r+1] != ' ' {
			r++
			if r+1 >= len(p) {
				r = len(p) - 1
				break
			}
		}

		i = r + 1
		for l < r {
			p[l], p[r] = p[r], p[l]
			l++
			r--
		}
	}
	return string(p)
}

// func reverseWords(s string) string {
// 	b := []byte(s)
// 	b = trim(b)
// 	b = reverse(b)
// 	b = findWordAndReverse(b)
// 	return string(b)
// }

// const SPACE = byte(rune(' '))

// func trim(s []byte) []byte {
// 	//trim start space
// 	for s[0] == SPACE {
// 		for i := 1; i < len(s); i++ {
// 			s[i-1] = s[i]
// 		}
// 		s = s[0 : len(s)-1]
// 	}

// 	//trim mid space

// 	l := 1
// 	swap := 0
// 	for i := 1; i < len(s); i++ {
// 		if s[i] == SPACE && s[i-1] == SPACE {
// 			swap++
// 		} else {
// 			if i != l {
// 				s[l] = s[i]
// 			}
// 			l++
// 		}
// 	}
// 	s = s[0 : len(s)-swap]

// 	//trim end space
// 	l = len(s) - 1
// 	for ; l > 0; l-- {
// 		if s[l] != SPACE {
// 			break
// 		}
// 	}

// 	return s[0 : l+1]
// }

// func reverse(s []byte) []byte {
// 	for i := 0; i < len(s)-i; i++ {
// 		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
// 	}
// 	return s
// }

//	func findWordAndReverse(s []byte) []byte {
//		prev := 0
//		for i := 1; i < len(s)+1; i++ {
//			if i == len(s) || s[i] == SPACE {
//				reverse(s[prev:i])
//				prev = i + 1
//			}
//		}
//		return s
//	}
func main() {
	quet := "a good   example"
	reverseWords(quet)
}
