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
	"math"
)

// Matrix Ai has dimension P[i-1] x P[i] for i = 1..n
func matrixChainOrder(P []int, i, j int) int {
	if i == j {
		return 0
	}

	min := math.MaxInt32

	// place parenthesis at different places between first
	// and last matrix, recursively calculate count of
	// multiplications for each parenthesis placement and
	// return the minimum count
	for k := i; k < j; k++ {
		count := matrixChainOrder(P, i, k) + matrixChainOrder(P, k+1, j) + P[i-1]*P[k]*P[j]

		if count < min {
			min = count
		}
	}

	// Return minimum count
	return min
}

func main() {
	P := []int{1, 2, 3, 4, 3}
	fmt.Println(matrixChainOrder(P, 1, len(P)-1))
}
