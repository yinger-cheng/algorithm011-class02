package main

import "fmt"

func main(){
	fmt.Println(maxProfit([]int{1,2,3,4,5,9}))
}

func maxProfit(prices []int) int {
	max := 0
	for i := 1 ; i < len(prices) ; i++ {
		if prices[i] > prices[i-1] {
			max = max + prices[i] - prices[i-1]
		}
	}
	return max
}
