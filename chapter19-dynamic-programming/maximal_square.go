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

func maximalSquare(B [][]byte) int {
	// row column size
	rows, cols, maxSquareLen := len(B), len(B[0]), 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if B[i][j] == 1 {
				sqlen, flag := 1, true
				for sqlen+i < rows && sqlen+j < cols && flag {
					for k := j; k <= sqlen+j; k++ {
						if B[i+sqlen][k] == 0 {
							flag = false
							break
						}
					}
					for k := i; k <= sqlen+i; k++ {
						if B[k][j+sqlen] == 0 {
							flag = false
							break
						}
					}
					if flag {
						sqlen++
					}
				}
				if maxSquareLen < sqlen {
					maxSquareLen = sqlen
				}
			}
		}
	}
	return maxSquareLen * maxSquareLen
}

func maximalSquareDP(B [][]byte) int {
	if len(B) == 0 || len(B[0]) == 0 {
		return 0
	}
	// row column size
	rows, cols := len(B), len(B[0])
	maxSquareLen := 0
	T := make([][]int, rows+1)
	for i := range T {
		T[i] = make([]int, cols+1)
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if B[i][j] == 1 {
				T[i][j] = min(T[i-1][j-1], min(T[i][j-1], T[i-1][j])) + 1
			}
			maxSquareLen = max(maxSquareLen, T[i][j])
		}
	}
	return maxSquareLen * maxSquareLen
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	B := [][]byte{{0, 1, 1, 0, 1},
		{1, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0}}
	fmt.Println(maximalSquare(B))
	fmt.Println(maximalSquareDP(B))
}
