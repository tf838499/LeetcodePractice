package main

import (
	"fmt"
	// "strconv"
	"sort"
)

/*
Medium
Monotonic Stack

doing

Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current day.

The span of the stock's price today is defined as the maximum number of consecutive days (starting from today and going backward)
for which the stock price was less than or equal to today's price.

For example, if the price of a stock over the next 7 days were [100,80,60,70,60,75,85], then the stock spans would be [1,1,1,2,1,4,6].

Implement the StockSpanner class:

StockSpanner() Initializes the object of the class.
int next(int price) Returns the span of the stock's price given that today's price is price.

Input
["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]
Output
[null, 1, 1, 1, 2, 1, 4, 6]
Explanation
StockSpanner stockSpanner = new StockSpanner();
stockSpanner.next(100); // return 1
stockSpanner.next(80);  // return 1
stockSpanner.next(60);  // return 1
stockSpanner.next(70);  // return 2
stockSpanner.next(60);  // return 1
stockSpanner.next(75);  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
stockSpanner.next(85);  // return 6
[[],[29],[91],[62],[76],[51]]

[null,1,2,1,2,1]
*/

// type StockSpanner struct {
// 	price []int
// 	span  []int
// 	max   int
// }

// func Constructor() StockSpanner {
// 	p := new(StockSpanner)
// 	return *p
// }

// func (this *StockSpanner) Next(price int) int {
// 	if this.price == nil {
// 		this.price = append(this.price, price)
// 		this.span = append(this.span, 1)
// 		this.max = 0
// 		return 1
// 	}

// 	ans := 1
// 	lenPrice := len(this.price) - 1
// 	for lenPrice >= 0 && this.price[lenPrice] <= price {
// 		ans = this.span[lenPrice] + ans
// 		lenPrice = lenPrice - this.span[lenPrice]
// 	}

// 	this.price = append(this.price, price)
// 	this.span = append(this.span, ans)
// 	return ans
// }

const MaxInt = int(^uint(0) >> 1)

type StockSpanner struct {
	list  []int // kept prices, descending order
	cusum []int // the cumulative sum of the number of consecutive days
}

func Constructor() StockSpanner {
	return StockSpanner{
		list:  []int{MaxInt}, // dummy to simplify
		cusum: []int{0},      // dummy to simplify
	}
}

func (s *StockSpanner) Next(price int) int {
	index := sort.Search(len(s.list), func(i int) bool {
		return price >= s.list[i]
	})

	s.list = append(s.list[:index], price)
	fmt.Println(s.cusum[:index])
	fmt.Println(s.cusum[len(s.cusum)-1] + 1)
	s.cusum = append(s.cusum[:index], s.cusum[len(s.cusum)-1]+1)
	return s.cusum[len(s.cusum)-1] - s.cusum[len(s.cusum)-2]
}

func main() {
	// [[],[28],[14],[28],[35],[46],[53],[66],[80],[87],[88]]
	// [null,1,1,2,4,5,6,7,8,9,10]
	// [null,1,1,3,4,5,6,7,8,9,10]
	s := Constructor()
	s.Next(28)
	s.Next(14)
	s.Next(28)
	s.Next(35)
	s.Next(46)
	s.Next(53)
	s.Next(66)
	s.Next(80)
	// s.Next(29)
	// s.Next(91)
	// s.Next(62)
	// s.Next(76)
	// s.Next(51)
	fmt.Println(s)
}
