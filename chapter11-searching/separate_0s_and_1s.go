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

// Function to segregate 0s and 1s
func separate0sand1s1(A []int) []int {
	count := 0 // Counts the no of zeros in arr

	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			count++
		}
	}

	// Loop fills the arr with 0 until count
	for i := 0; i < count; i++ {
		A[i] = 0
	}
	// Loop fills remaining arr space with 1
	for i := count; i < len(A); i++ {
		A[i] = 1
	}
	return A
}

// Function to put all 0s on left and all 1s on right
func separate0sand1s2(A []int) []int {
	// Initialize left and right indexes
	left, right := 0, len(A)-1
	for left < right {
		// Increment left index while we see 0 at left
		for A[left] == 0 && left < right {
			left++
		}

		// Decrement right index while we see 1 at right
		for A[right] == 1 && left < right {
			right--
		}

		// If left is smaller than right then there is a 1 at left
		// and a 0 at right.  Swap A[left] and A[right]
		if left < right {
			A[left] = 0
			A[right] = 1
			left++
			right--
		}
	}
	return A
}

func main() {
	A := []int{1, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1}
	fmt.Println(separate0sand1s1(A))
	fmt.Println(separate0sand1s2(A))
}
