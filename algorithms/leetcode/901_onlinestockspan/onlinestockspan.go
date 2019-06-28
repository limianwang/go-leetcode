package onlinestockspan

// StockSpanner
type StockSpanner struct {
	stocks  []int
	weights []int
}

func (this *StockSpanner) Next(price int) int {
	weight := 1

	for len(this.stocks) > 0 && this.stocks[len(this.stocks)-1] <= price {
		weight += this.weights[len(this.weights)-1]
		this.stocks = this.stocks[:len(this.stocks)-1]
		this.weights = this.weights[:len(this.weights)-1]
	}

	this.stocks = append(this.stocks, price)
	this.weights = append(this.weights, weight)

	return weight
}

func Constructor() StockSpanner {
	return StockSpanner{
		stocks:  []int{},
		weights: []int{},
	}
}
