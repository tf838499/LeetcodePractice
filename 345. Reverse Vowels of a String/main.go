package main

import (
	"fmt"
	// "strconv"
)

/*
easy
string
two pointer
done

Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

Input: s = "hello"
Output: "holle"

Input: s = "leetcode"
Output: "leotcede"
*/
func reverseVowels(s string) string {
	// vowels := "ae"
	vowels := []rune("aeiouAEIOU")
	a := []rune(s)
	leftPoint := 0
	rightsPoint := len(a) - 1
	LeftCheck := true
	RightCheck := true
	for leftPoint < rightsPoint {

		for _, i := range vowels {

			if i == a[leftPoint] && LeftCheck {
				LeftCheck = false
			}
			if i == a[rightsPoint] && RightCheck {
				RightCheck = false
			}
			if RightCheck == false && LeftCheck == false {
				a[leftPoint], a[rightsPoint] = a[rightsPoint], a[leftPoint]
				RightCheck, LeftCheck = true, true
				break
			}
		}
		if RightCheck {
			rightsPoint--
		}
		if LeftCheck {
			leftPoint++
		}

		// a[leftPoint], a[rightsPoint] = a[rightsPoint], a[leftPoint]

	}
	// fmt.Println(string(a))
	return string(a)
}

func main() {
	words := "Euston saw I was not Sue."
	// words := []byte{"H", "a", "n", "n", "a", "h"}

	// words := ["bella","label","roller"]
	reverseVowels(words)

}
