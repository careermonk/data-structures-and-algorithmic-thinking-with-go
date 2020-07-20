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

func Mins(value int, values ...int) int {
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

func minDistance(A string, B string) int {
	T := make([][]int, len(A)+1)
	for i := range T {
		T[i] = make([]int, len(B)+1)
		T[i][0] = i
	}
	for j := range T[0] {
		T[0][j] = j
	}
	for i := 1; i < len(T); i++ {
		for j := 1; j < len(T[0]); j++ {
			offset := 1
			if A[i-1] == B[j-1] {
				offset = 0
			}
			T[i][j] = Mins(T[i-1][j]+1, T[i][j-1]+1, T[i-1][j-1]+offset)
		}
	}
	return T[len(T)-1][len(T[0])-1]
}

func main() {
	A := "intention"
	B := "execution"
	fmt.Println("Edit distance between ", A, " and ", B, " is ", minDistance(A, B))

	A = "horse"
	B = "ros"
	fmt.Println("Edit distance between ", A, " and ", B, " is ", minDistance(A, B))
}
