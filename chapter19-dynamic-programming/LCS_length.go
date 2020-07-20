package main

import (
	"fmt"
)

func LCSLength(X, Y string) int {
	m, n := len(X), len(Y)
	LCS := make([][]int, m+1)
	for i := range LCS {
		LCS[i] = make([]int, n+1) // defaults to 0
	}
	fmt.Println(LCS)
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
}
