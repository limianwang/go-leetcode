/*
Leetcode #901
*/

package onlinestockspan

import (
	"math"
)

type StockSpanner struct {
	stocks      []int
	count       int
	lastHighest int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stocks:      []int{},
		count:       0,
		lastHighest: math.MaxInt64,
	}
}

func (this *StockSpanner) Next(price int) int {
	days := 0
	this.stocks = append(this.stocks, price)

	for i := len(this.stocks) - 1; i >= 0; i-- {
		if price >= this.stocks[i] {
			days++
		} else {
			break
		}
	}
	return days
}
