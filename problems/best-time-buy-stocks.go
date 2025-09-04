package problems

import "fmt"

//input prices = [7,1,5,3,6,4]

// brute force
func MaxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}
	return profit
}

func MaxProfit2(prices []int) int {
	profit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	fmt.Println(profit)
	return profit
}
