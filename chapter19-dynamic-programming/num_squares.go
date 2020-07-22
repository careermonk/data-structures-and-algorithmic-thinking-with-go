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
	"math"
)

func numSquares(n int) int {
	T := make([]int, n+1)
	T[0], T[1] = 0, 1

	for i := 2; i <= n; i++ {
		T[i] = math.MaxInt32
		for j := 1; i-j*j >= 0; j++ {
			T[i] = min(T[i], T[i-j*j]+1)
		}
	}

	return T[n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
}
