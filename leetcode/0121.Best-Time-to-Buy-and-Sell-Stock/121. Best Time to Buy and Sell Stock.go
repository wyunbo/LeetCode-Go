package leetcode

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit, min := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if profit < prices[i]-min {
			profit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return profit
}
