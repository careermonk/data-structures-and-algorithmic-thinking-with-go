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

func matrixChainOrderR(P []int, i, j int) int {
	if i == j {
		return 0
	}
	const MaxUint = ^uint(0)
	const INT_MAX = int(MaxUint >> 1)
	min := INT_MAX

	for k := i; k < j; k++ {
		count := matrixChainOrderR(P, i, k) + matrixChainOrderR(P, k+1, j) + P[i-1]*P[k]*P[j]
		if count < min {
			min = count
		}
	}

	// Return minimum count
	return min
}

func matrixChainOrder(P []int) int {
	n := len(P)
	M := make([][]int, n)
	for i := range M {
		M[i] = make([]int, n)
	}
	// cost is zero when multiplying one matrix.
	for i := 1; i < n; i++ {
		M[i][i] = 0
	}
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	// L is chain length.
	for L := 2; L < n; L++ {
		for i := 1; i < n-L+1; i++ {
			j := i + L - 1
			M[i][j] = MaxInt
			for k := i; k <= j-1; k++ {
				// q = cost/scalar multiplications
				q := M[i][k] + M[k+1][j] + P[i-1]*P[k]*P[j]
				if q < M[i][j] {
					M[i][j] = q
				}
			}
		}
	}
	return M[1][n-1]
}

func main() {
	P := []int{1, 2, 3, 4}
	fmt.Println(matrixChainOrderR(P, 1, len(P)-1))
	fmt.Println(matrixChainOrder([]int{1, 2, 3, 4}))
}
