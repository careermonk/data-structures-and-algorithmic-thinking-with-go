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

func DutchFlagProblem(A []int) []int {
	low, mid, high := 0, 0, len(A)-1
	for mid <= high {
		switch A[mid] {
		case 0:
			A[low], A[mid] = A[mid], A[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			A[mid], A[high] = A[high], A[mid]
			high--
		}
	}
	return A
}

func main() {
	A := []int{1, 0, 2, 0, 1, 2, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1}
	fmt.Println(DutchFlagProblem(A))
	fmt.Println(DutchFlagProblem(A))
}
