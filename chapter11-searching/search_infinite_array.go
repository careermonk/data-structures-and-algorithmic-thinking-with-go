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

func search_infinite_array1(A []int, K int) {
	left, right := 0, C
	for A[right] < K {
		left = right
		right = right + C
	}
	return binarySearch(A, left, right, K)
}

func search_infinite_array2(A []int, K int) {
	left, right := 0, 1
	for A[right] < K {
		left = right
		right = 2 * right
	}
	return binarySearch(A, left, right, K)
}
