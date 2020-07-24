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

func firstOccuranceR(A []int, low, high, data int) int {
	if high >= low {
		mid := low + (high-low)/2
		if (mid == 0 || data > A[mid-1]) && A[mid] == data {
			return mid
		} else if data > A[mid] {
			return firstOccuranceR(A, (mid + 1), high, data)
		} else {
			return firstOccuranceR(A, low, (mid - 1), data)
		}
	}
	return -1
}

func firstOccurance(A []int, data int) int {
	low, high, res := 0, len(A)-1, -1
	for low <= high {
		// Normal Binary Search Logic
		mid := (low + high) / 2
		if A[mid] > data {
			high = mid - 1
		} else if A[mid] < data {
			low = mid + 1
		} else { // If A[mid] is same as data, we update res and move to the left half
			res = mid
			high = mid - 1
		}
	}
	return res
}

func main() {
	A := []int{1, 2, 2, 2, 2, 3, 4, 7, 8, 8}
	fmt.Println(firstOccuranceR(A, 0, len(A)-1, 2))
	fmt.Println(firstOccuranceR(A, 0, len(A)-1, 8))
	fmt.Println(firstOccurance(A, 2))
	fmt.Println(firstOccurance(A, 8))
}
