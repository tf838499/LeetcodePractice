package main

// "strconv"
// "strings"
// "math"

/*
You are given an integer array prices where prices[i] is
the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock.
You can only hold at most one share of the stock at any time.
However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.

Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5),
 profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6),
 profit = 6-3 = 3.
Total profit is 4 + 3 = 7.
*/
func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	Profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			Profit += (prices[i] - prices[i-1])
		}
	}
	return Profit
}
func Accumulate(elements ...int) int {
	sum := 0
	for _, elem := range elements {
		sum += elem
	}
	return sum
}

func maxProfit(prices []int) int {

	profit := make([]int, 0)

	for i := 0; i < len(prices)-1; i++ {

		if prices[i] < prices[i+1] {
			profit = append(profit, (prices[i+1] - prices[i]))
		}
	}

	return Accumulate(profit...)
}
func main() {
	// isPalindrome("A man, a plan, a canal: Panama")
	prices := []int{7, 6, 4, 3, 1}
	maxProfit(prices)
}
