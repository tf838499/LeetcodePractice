package main

import (
	"fmt"
	// "strconv"
)

/*
easy
hashmap
done

Given a string array words, return an array of all characters
that show up in all strings within the words (including duplicates).
You may return the answer in any order.

Input: words = ["bella","label","roller"]
Output: ["e","l","l"]
Input: words = ["cool","lock","cook"]
Output: ["c","o"]
*/
func commonChars(words []string) []string {

	// n := make(map[int32]int)

	n := make([]map[int32]int, len(words))
	for i, e := range words {

		b := []rune(e)
		t := make(map[int32]int)
		for _, p := range b {
			t[p] = t[p] + 1
		}
		n[i] = t
	}
	t := n[0]
	for u, o := range t {
		min := o
		for i := 0; i < len(words); i++ {
			fmt.Println(n[i][u])
			if min > n[i][u] {
				fmt.Println(n[i][u])
				min = n[i][u]
			}
		}

		fmt.Println(u)
		fmt.Println(o)
	}
	// fmt.Println(patu)
	return words
}

func main() {
	words := []string{
		"bella",
		"label",
		"roller",
	}
	// words := ["bella","label","roller"]
	commonChars(words)

}
