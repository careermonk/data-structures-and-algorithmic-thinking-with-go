// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

package main

import (
	"fmt"
)

func coinChange(coins []int, C int) int {
	T := make([]int, C+1)
	const MaxUint = ^uint(0)
	const INT_MAX = int(MaxUint >> 1)
	for i := 1; i <= C; i++ {
		minCoin := INT_MAX
		for _, coin := range coins {
			if i-coin >= 0 && T[i-coin] != -1 {
				minCoin = min(minCoin, T[i-coin]+1)
			}
		}
		if minCoin == INT_MAX {
			T[i] = -1
		} else {
			T[i] = minCoin
		}
	}
	return T[C]
}

func coinChangeR(coins []int, C int) int {
	// base case
	if C == 0 {
		return 0
	}

	const MaxUint = ^uint(0)
	const INT_MAX = int(MaxUint >> 1)
	result := INT_MAX

	// Try every coin that has smaller value than V
	for i := 0; i < len(coins); i++ {
		if coins[i] <= C {
			subRes := coinChangeR(coins, C-coins[i])
			// Check for INT_MAX to avoid overflow and see if result can minimized
			if subRes != INT_MAX && subRes+1 < result {
				result = subRes + 1
			}
		}
	}
	return result
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	coins := []int{1, 2, 5}
	fmt.Println(coinChange(coins, 11))

	coins = []int{9, 6, 5, 1}
	fmt.Println(coinChange(coins, 11))

	coins = []int{1, 2, 5}
	fmt.Println(coinChangeR(coins, 11))

	coins = []int{9, 6, 5, 1}
	fmt.Println(coinChangeR(coins, 11))
}
