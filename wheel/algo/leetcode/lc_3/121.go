package lc_3

func maxProfitQ1(prices []int) int {
	curMin, curMax := 10001, 0
	for _, price := range prices {
		if price > curMin {
			curMax = max(curMax, price-curMin)
		}
		curMin = min(price, curMin)
	}
	return curMax
}
