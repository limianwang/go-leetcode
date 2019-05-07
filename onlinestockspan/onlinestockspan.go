/*
Leetcode #901
*/

package onlinestockspan

import (
	"github.com/limianwang/goexamples/datastructures/stack"
)

type StockSpanner struct {
	stocks      *stack.Stack
	weight      *stack.Stack
	count       int
	lastHighest int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stocks: stack.New(),
		weight: stack.New(),
	}
}

func (this *StockSpanner) Next(price int) int {
	w := 1
	for !this.stocks.IsEmpty() && this.stocks.Peek().(int) <= price {
		this.stocks.Pop()
		w += this.weight.Pop().(int)
	}
	this.stocks.Push(price)
	this.weight.Push(w)
	return w
}
