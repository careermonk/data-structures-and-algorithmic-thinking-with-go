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

func longestPalindromeSubstring(A string) int {
	n := len(A)
	L := make([][]bool, n)
	for i := range L {
		L[i] = make([]bool, n) // defaults to false
	}
	max := 1
	for i := 0; i < n-1; i++ {
		L[i][i] = true
		if A[i] == A[i+1] {
			L[i][i+1] = true
			max = 2
		}
	}
	for k := 3; k <= n; k++ {
		for i := 1; i < n-k+1; i++ {
			j := i + k - 1
			if A[i] == A[j] && L[i+1][j-1] {
				L[i][j] = true
				max = k
			} else {
				L[i][j] = false
			}
		}
	}
	return max
}

func main() {
	fmt.Print(longestPalindromeSubstring("babad"))

}
