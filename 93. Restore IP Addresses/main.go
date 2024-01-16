package main

import (
	"strconv"
	"strings"
)

// "fmt"
// "strconv"

/*
A valid IP address consists of exactly four integers separated
by single dots. Each integer is between 0 and 255 (inclusive) and
cannot have leading zeros.

For example,
"0.1.2.201" and "192.168.1.1" are valid IP addresses,
but "0.011.255.245", "192.168.1.312" and
"192.168@1.1" are invalid IP addresses.
Given a string s containing only digits,
return all possible valid IP addresses
that can be formed by inserting dots into s.
You are not allowed to reorder or remove any digits in s.
You may return the valid IP addresses in any order.

Example 1:
Input: s = "25525511135"
Output: ["255.255.11.135","255.255.111.35"]

Example 2:
Input: s = "0000"
Output: ["0.0.0.0"]

Example 3:
Input: s = "101023"
Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
*/
// func restoreIpAddresses(s string) []string {
// 	result := []string{}
// 	var path []string
// 	var dfs func(s string, ind int)
// 	dfs = func(s string, ind int) {
// 		value := 0
// 		if len(path) >= 4 {
// 			if ind == len(s) {
// 				str := strings.Join(path, ".")
// 				result = append(result, str)
// 			}
// 			return
// 		}
// 		for i := ind; i < len(s); i++ {
// 			if i != ind && s[ind] == '0' { // 含有前导 0，无效
// 				break
// 			}
// 			candidate := s[ind : i+1]
// 			// s[ind : i+1]
// 			value, _ = strconv.Atoi(candidate)

// 			if value >= 0 && value <= 255 {
// 				path = append(path, candidate) // 符合条件的就进入下一层
// 				dfs(s, i+1)
// 				path = path[:len(path)-1]
// 			} else { // 如果不满足条件，再往后也不可能满足条件，直接退出
// 				break
// 			}
// 			// dfs()
// 		}
// 	}
// 	dfs(s, 0)

// 	// for i := 0; i < len(s); i++ {
// 	// 	candidate := s[:i+1]

// 	// }
// 	return result
// }

func restoreIpAddresses(s string) []string {
	ans := []string{}
	var path []string
	var substring func(str string, strind int)
	substring = func(str string, strind int) {
		if len(path) >= 4 {
			if len(str) == 0 {
				joinstr := strings.Join(path, ".")
				ans = append(ans, joinstr)
			}
			return
		}

		for i := 0; i < len(str); i++ {
			if str[strind] == '0' && i != strind { // 含有前导 0，无效
				break
			}
			candidate := str[:i+1]
			value, _ := strconv.Atoi(candidate)
			if value >= 0 && value <= 255 {
				path = append(path, candidate)
				substring(str[i+1:], 0)
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}
	substring(s, 0)
	return ans
}

func main() {

	restoreIpAddresses("0000")

}

// 255 3
// 2 2
//
//
//
//
