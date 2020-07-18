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

func maxRepetitions1(A []int) []int {
	counter, maxCounter, maxElement := 0, 0, -1
	for i := 0; i < len(A); i++ {
		counter = 0
		for j := 0; j < len(A); j++ {
			if A[i] == A[j] {
				counter++
			}
		}
		if counter > maxCounter {
			maxCounter = counter
			maxElement = A[i]
		}
	}
	return []int{maxCounter, maxElement}
}

func maxRepititions2(A []int) []int {
	maxCounter, maxElement, n := 0, -1, len(A)
	for i := 0; i < n; i++ {
		A[A[i]%n] += n
	}
	for i := 0; i < n; i++ {
		if A[i]/n > maxCounter {
			maxCounter = A[i] / n
			maxElement = i
		}
	}
	return []int{maxCounter, maxElement}
}

func main() {
	A := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(A)

	// Test the nested loop method.
	fmt.Println(maxRepititionsBruteForce2(A))
}
