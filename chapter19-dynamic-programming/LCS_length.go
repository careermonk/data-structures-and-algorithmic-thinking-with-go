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

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// Initial Call: LCSLengthRecursive(X, 0, m-1, Y, 0, n-1);
func LCSLengthRecursive(X string, i, m int, Y string, j, n int) int {
	if i == m || j == n {
		return 0
	} else if X[i] == Y[j] {
		return 1 + LCSLengthRecursive(X, i+1, m, Y, j+1, n)
	} else {
		return max(LCSLengthRecursive(X, i+1, m, Y, j, n), LCSLengthRecursive(X, i, m, Y, j+1, n))
	}
}

func LCSLength2(X, Y string) int {
	m, n := len(X), len(Y)
	return LCSLengthRecursive(X, 0, m, Y, 0, n)
}

func LCSLength(X, Y string) int {
	m, n := len(X), len(Y)
	LCS := make([][]int, m+1)
	for i := range LCS {
		LCS[i] = make([]int, n+1) // defaults to 0
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			LCS[i][j] = LCS[i+1][j+1] // matching X[i] to Y[j]

			if X[i] == Y[j] {
				LCS[i][j]++ // we get a matching pair
			}

			// the other two cases – inserting a gap
			if LCS[i][j+1] > LCS[i][j] {
				LCS[i][j] = LCS[i][j+1]
			}

			if LCS[i+1][j] > LCS[i][j] {
				LCS[i][j] = LCS[i+1][j]
			}
		}
	}
	return LCS[0][0]
}
func main() {
	fmt.Println(LCSLength("abcde", "ace"))
	fmt.Println(LCSLength("abc", "def"))

	fmt.Println(LCSLength2("abcde", "ace"))
	fmt.Println(LCSLength2("abc", "def"))
}
