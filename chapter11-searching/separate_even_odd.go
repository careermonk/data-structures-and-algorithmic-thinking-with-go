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

func separateEvenOdd(A []int) []int {
	left, right := 0, len(A)-1
	for left < right {
		// Increment left index while we see 0 at left
		for A[left]%2 == 0 && left < right {
			left++
		}
		// Decrement right index while we see 1 at right
		for A[right]%2 == 1 && left < right {
			right--
		}
		if left < right {
			// Swap A[left] and A[right]
			A[left], A[right] = A[right], A[left]
			left++
			right--
		}
	}
	return A
}

func main() {
	P := []int{12,34,45,9,8,90,3}
	P=separateEvenOdd(P)
	fmt.Println(P)
}
