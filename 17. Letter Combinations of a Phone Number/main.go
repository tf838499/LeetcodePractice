package main

// "strconv"

/*
Given a string containing digits from 2-9 inclusive,
return all possible letter combinations that the number could represent.
Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons)
is given below. Note that 1 does not map to any letters.

Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Example 2:

Input: digits = ""
Output: []
Example 3:

Input: digits = "2"
Output: ["a","b","c"]
2:abc
3:def
4:ghi
5:jkl
6:mno
7:pqrs
8:tuv
9:wxyz
*/
// func letterCombinations(digits string) []string {
// 	phoneNumMap := map[byte]string{
// 		'2': "abc",
// 		'3': "def",
// 		'4': "ghi",
// 		'5': "jkl",
// 		'6': "mno",
// 		'7': "pqrs",
// 		'8': "tuv",
// 		'9': "wxyz",
// 	}
// 	var result []string
// 	if len(digits) == 0 {
// 		return result
// 	}
// 	// comb :=
// 	var comb func(index int, p string, n []byte)
// 	comb = func(index int, p string, n []byte) {
// 		if len(p) == 0 {
// 			dst := make([]byte, len(digits))
// 			copy(dst, n)
// 			result = append(result, string(dst))
// 			return
// 		}

// 		letter := phoneNumMap[p[0]]
// 		for i := range phoneNumMap[p[0]] {
// 			n = append(n, letter[i])
// 			comb(index+1, p[1:], n)
// 			n = n[0 : len(n)-1]
// 		}

// 	}
// 	comb(0, digits, []byte{})
// 	return result
// }
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	phoneMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var output []string

	var backtrack func(combination string, nextDigits string)
	backtrack = func(combination string, nextDigits string) {
		if nextDigits == "" {
			output = append(output, combination)
		} else {
			letters := phoneMap[nextDigits[0]-'2']
			for _, letter := range letters {
				backtrack(combination+string(letter), nextDigits[1:])
			}
		}
	}

	backtrack("", digits)
	return output
}
func main() {
	// words := []string{
	// 	"bella",
	// 	"label",
	// 	"roller",
	// }
	// words := ["bella","label","roller"]
	letterCombinations("23")

}
