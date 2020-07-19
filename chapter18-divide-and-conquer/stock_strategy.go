// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

package main

import "fmt"

func minIndex(prices []int, i, j int) int {
	result := i
	for k := i + 1; k <= j; k++ {
		if prices[k] < prices[result] {
			result = k
		}
	}
	return result
}

func maxIndex(prices []int, i, j int) int {
	result := i
	for k := i + 1; k <= j; k++ {
		if prices[k] > prices[result] {
			result = k
		}
	}
	return result
}

func stockStrategy1(prices []int) []int {
	buyDateIndex, sellDateIndex, profit, n := 0, 0, 0, len(prices)
	for i := 0; i < n; i++ { // indicates buy date
		for j := i + 1; j < n; j++ { // indicates sell date
			if prices[j]-prices[i] > profit {
				profit = prices[j] - prices[i]
				buyDateIndex = i
				sellDateIndex = j
			}
		}
	}
	return []int{profit, buyDateIndex, sellDateIndex}
}

func stockStrategy(prices []int, left, right int) (int, int, int) {
	// If the array has just one element, we return that the profit is zero
	// but the minimum and maximum values are just that array value.
	if left == right {
		return 0, left, right

	}
	mid := left + (right-left)/2

	leftProfit, leftBuyDateIndex, leftSellDateIndex := stockStrategy(prices, left, mid)
	rightProfit, rightBuyDateIndex, rightSellDateIndex := stockStrategy(prices, mid+1, right)

	minIndexLeft := minIndex(prices, left, mid)
	maxIndexRight := maxIndex(prices, mid+1, right)

	centerProfit := prices[maxIndexRight] - prices[minIndexLeft]
	if (centerProfit > leftProfit) && (centerProfit > rightProfit) {
		return centerProfit, minIndexLeft, maxIndexRight
	} else if (leftProfit > centerProfit) && (leftProfit > rightProfit) {
		return leftProfit, leftBuyDateIndex, leftSellDateIndex
	} else {
		return rightProfit, rightBuyDateIndex, rightSellDateIndex
	}
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(stockStrategy1(prices))
	fmt.Println(stockStrategy(prices, 0, len(prices)-1))
}
