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

func sortedSquares_naive(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}
	sort.Ints(A)
	return A
}

func sortedSquares(A []int) []int {
	result := make([]int, len(A))
	left, right := 0, len(A)-1
	for i := len(A) - 1; i >= 0; i-- {
		if abs(A[left]) < abs(A[right]) {
			result[i] = A[right] * A[right]
			right--
		} else {
			result[i] = A[left] * A[left]
			left++
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// Driver program to test above functions
func main() {
	fmt.Println(sortedSquares_naive([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
