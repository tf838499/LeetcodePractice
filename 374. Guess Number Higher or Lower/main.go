package main

import (
	"fmt"
	// "strconv"
)

/*
easy
binary search
day 1
筆記: 需要在多注意題目意思，以及給予的線索，在binary search時要怎麼制定左右邊界很重要
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API int guess(int num), which returns three possible results:

-1: Your guess is higher than the number I picked (i.e. num > pick).
1: Your guess is lower than the number I picked (i.e. num < pick).
0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.

Input: n = 10, pick = 6
Output: 6

Input: n = 1, pick = 1
Output: 1

Input: n = 2, pick = 1
Output: 1

*/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	h := n
	l := 0

	for true {
		mid := (h-l)/2 + l
		ans := guess(mid)
		if ans == 0 {
			return mid
		}
		if ans == -1 {
			h = mid
		} else {
			l = mid
		}

	}
	return -1
}
func guessNumber(n int) int {
	left, right := 1, n
	for left <= right {
		mid := (left + right) / 2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func main() {
	words := "Hannah"
	// words := []byte{"H", "a", "n", "n", "a", "h"}

	// words := ["bella","label","roller"]
	reverseString([]byte(words))

}
