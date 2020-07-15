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
	"sort"
)

func binarySearch(data int, A []int) bool {
	low := 0
	high := len(A) - 1
	for low <= high {
		mid := (low + high) / 2
		if A[mid] < data {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low == len(A) || A[low] != data {
		return false
	}
	return true
}

func find(A, B []int, K int) bool {
	sort.Ints(A)                  // O(nlogn)
	for i := 0; i < len(B); i++ { // O(n)
		c := K - B[i]           // O(1)
		if binarySearch(c, A) { // O(logn)
			return true
		}
	}
	return false
}

// Driver code
func main() {
	A := []int{11, 2, 9, 7, 11, 5, 73, 80, 10} // should be a sorted array
	B := []int{1, 2, 4, 10, 11, 4, 3, 8, 30}   // should be a sorted array
	fmt.Println(find(A, B, 5))
	fmt.Println(find(A, B, 15))
	fmt.Println(find(A, B, 103))
}
