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
func matrixChainOrder(P []int) int {
	// For simplicity of the program, one extra row and one extra column are allocated in M[][].  
	// 0th row and 0th column of M[][] are not used
	n := len(P)
	M := make([][]int, n)
	for i := range M {
		M[i] = make([]int, n) // defaults to fa0lse
	}

	// M[i,j] = Minimum number of scalar multiplications needed to compute the 
	// matrix A[i]A[i+1]...A[j] = A[i..j] where dimension of A[i] is P[i-1] x P[i]
	// cost is zero when multiplying one matrix.
	for i := 1; i < n; i++ {
		M[i][i] = 0
	}

	// L is chain length.
	for L := 2; L < n; L++ {
		for i := 1; i < n-L+1; i++ {
			j := i + L - 1
			M[i][j] = math.MaxInt32
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
	P := []int{1, 2, 3, 4, 3}
	fmt.Println(matrixChainOrder(P))
}
