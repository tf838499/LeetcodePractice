package main

import (
// "sort"
// "strconv"
)

/*
Assume you are an awesome parent and want to give your children some cookies.
But, you should give each child at most one cookie.

Each child i has a greed factor g[i],
which is the minimum size of a cookie that the child will be content with;
and each cookie j has a size s[j]. If s[j] >= g[i],
we can assign the cookie j to the child i,
and the child i will be content.
Your goal is to maximize the number of your content children and output the maximum number.

Input: g = [1,2,3], s = [1,1]
Output: 1
Explanation: You have 3 children and 2 cookies. The greed factors of 3 children are 1, 2, 3.
And even though you have 2 cookies, since their size is both 1,
you could only make the child whose greed factor is 1 content.
You need to output 1.

Input: g = [1,2], s = [1,2,3]
Output: 2
Explanation: You have 2 children and 3 cookies. The greed factors of 2 children are 1, 2.
You have 3 cookies and their sizes are big enough to gratify all of the children,
You need to output 2.

*/
func findContentChildren(g []int, s []int) int {
	ans := 0
	Sort.Ints(g)
	Sort.Ints(s)
	for len(s) != 0 && len(g) != 0 {
		if g[0] < s[0] {
			s[0] = 0
			g[0] = 0
		} else {
			g[0] = g[0] - s[0]
			s[0] = 0
		}

		if g[0] == 0 {
			g = g[1:]
			ans++
		}
		if s[0] == 0 {
			s = s[1:]
		}
	}
	return ans
}

// func findContentChildren(g []int, s []int) int {
// 	sort.Ints(g)
// 	sort.Ints(s)
// 	chil, cook := 0, 0
// 	for chil < len(g) && cook < len(s) {
// 		if g[chil] <= s[cook] {
// 			chil++
// 		}
// 		cook++ //increase num of cookies until the 1st child is content
// 	}
// 	return chil
// }
func main() {
	g := []int{1, 2, 3}
	s := []int{1, 2, 2}
	findContentChildren(g, s)

}
