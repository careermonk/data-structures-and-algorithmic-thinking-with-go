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

func convertArraytoSawToothWave(A []int) {
	// Flag true indicates relation "<" is expected, else ">" is expected.  The first expected relation is "<"
	flag, n := true, len(A)

	for i := 0; i <= n-2; i++ {
		if flag {
			// If we have a situation like A > B > C, we get A > B < C by swapping B and C
			if A[i] > A[i+1] {
				A[i], A[i+1] = A[i+1], A[i]
			}
		} else {
			// If we have a situation like A < B < C, we get A < C > B by swapping B and C
			if A[i] < A[i+1] {
				A[i], A[i+1] = A[i+1], A[i]
			}
		}
		if flag {
			flag = false
		} else {
			flag = true
		}
	}
}

// Driver program to test above functions
func main() {
	A := []int{0, -6, 9, 13, 10, -1, 8, 12, 54, 14, -5}
	convertArraytoSawToothWave(A)
	fmt.Println(A)

	A = []int{4, 3, 7, 8, 6, 2, 1}
	convertArraytoSawToothWave(A)
	fmt.Println(A)
}
