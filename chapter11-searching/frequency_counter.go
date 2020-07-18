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

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func frequencyCounter(A []int) {
	pos, n := 0, len(A)
	for pos < n {
		expectedPos := A[pos] - 1
		if A[pos] > 0 && A[expectedPos] > 0 {
			A[pos], A[expectedPos] = A[expectedPos], A[pos]
			A[expectedPos] = -1
		} else if A[pos] > 0 {
			A[expectedPos]--
			A[pos] = 0
			pos++
		} else {
			pos++
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(i+1, " frequency is ", abs(A[i]))
	}
}

func main() {
	A := []int{10, 10, 9, 4, 7, 6, 5, 2, 3, 2, 1}
	fmt.Println(A)
	frequencyCounter(A)
}
