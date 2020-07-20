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

import (
	"fmt"
)

func calculate(V [][]int, i, j int) int {
	if i <= j {
		return V[i][j]
	}
	return 0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// Iterative function to maximize the number of coins collected by a player,
// assuming that opponent also plays optimally
func optimalStrategy(coins []int) int {
	n := len(coins)
	// base case: one pot left, only one choice possible
	if n == 1 {
		return coins[0]
	}
	// if we're left with only two pots, choose one with maximum coins
	if n == 2 {
		return max(coins[0], coins[1])
	}
	// create a dynamic 2D matrix to store sub-problem solutions
	V := make([][]int, n)
	for i := range V {
		V[i] = make([]int, n) // defaults to false
	}

	for iteration := 0; iteration < n; iteration++ {
		for i, j := 0, iteration; j < n; i, j = i+1, j+1 {
			start := coins[i] + min(calculate(V, i+2, j), calculate(V, i+1, j-1))
			end := coins[j] + min(calculate(V, i+1, j-1), calculate(V, i, j-2))
			V[i][j] = max(start, end)
		}
	}
	return V[0][n-1]
}

func main() {
	// pots of gold arranged in a line
	coins := []int{4, 6, 2, 3}
	fmt.Println("Maximum coins collected by player is ", optimalStrategy(coins))

	// pots of gold arranged in a line
	coins = []int{5, 3, 7, 10}
	fmt.Println("Maximum coins collected by player is ", optimalStrategy(coins))
}
