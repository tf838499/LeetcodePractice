package main

import (
	// "fmt"
	"strconv"
)

/*
easy
string
math
done

You are given a positive integer num consisting only of digits 6 and 9.

Return the maximum number you can get by changing at most one digit (6 becomes 9, and 9 becomes 6).

*/
func maximum69Number(num int) int {
	// runeNum := []rune(num)

	// s1 := strconv.FormatInt(int64(num), 10)
	// ori := []byte(strconv.Itoa(num))
	// ans := num
	maxAns := num

	for i := 0; i < len(string(num)); i++ {
		ans := []byte(strconv.Itoa(num))

		// fmt.Println(ans)
		if ans[i] == 57 {
			ans[i] = 54
		} else if ans[i] == 54 {
			ans[i] = 57
		}
		// fmt.Println(string(ans))
		ansStr := string(ans)
		// ansInt :=ansInt
		ansInt, _ := strconv.Atoi(ansStr)

		if ansInt > maxAns {
			maxAns = ansInt
		}
	}
	return maxAns
}

// func maximum69Number(num int) int {
// 	r, _ := strconv.Atoi((strings.Replace(strconv.Itoa(num), "6", "9", 1)))
// 	return r
// }
func main() {
	num := 9669
	maximum69Number(num)
}
