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

func findApplesCount(A [][]int) int {
	n, m := len(A), len(A[0])
	//computing the vertical prefix sum for columns
	S := make([][]int, n)
	for i := range S {
		S[i] = make([]int, m) // defaults to 0
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			S[i][j] = A[i][j]
			if S[i][j] < S[i][j]+S[i][j-1] {
				S[i][j] += S[i][j-1]
			}

			if S[i][j] < S[i][j]+S[i-1][j] {
				S[i][j] += S[i-1][j]
			}

		}
	}
	return S[n-1][m-1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	A := [][]int{{0, 2, 4, 1},
		{4, 8, 3, 7},
		{2, 3, 6, 2},
		{9, 7, 8, 3},
		{1, 5, 9, 4}}
	fmt.Println(findApplesCount(A))
}
