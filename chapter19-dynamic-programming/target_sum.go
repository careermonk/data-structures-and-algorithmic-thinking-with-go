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

func targtSum1(A []int, n, T int) bool {
	// Base Cases
	if T == 0 {
		return true
	}
	if n == 0 && T != 0 {
		return false
	}

	// If last element is greater than T, then ignore it
	if A[n-1] > T {
		return targtSum1(A, n-1, T)
	}
	return targtSum1(A, n-1, T) || targtSum1(A, n-1, T-A[n-1])
}

func targtSum(A []int, T int) bool {
	n := len(A)
	M := make([][]bool, n+1)
	for i := range M {
		M[i] = make([]bool, T+1) // defaults to false
	}
	// If T is 0, then answer is true
	for i := 0; i <= n; i++ {
		M[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= T; j++ {
			if j < A[i-1] {
				M[i][j] = M[i-1][j]
			}
			if j >= A[i-1] {
				M[i][j] = M[i-1][j] || M[i-1][j-A[i-1]]
			}
		}
	}
	return M[n][T]
}

func main() {
	A := []int{3, 34, 4, 12, 5, 2}
	fmt.Println(targtSum1(A, len(A), 9))
	fmt.Println(targtSum(A, 9))
}
