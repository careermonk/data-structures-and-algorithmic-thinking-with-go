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

// O(n^3)
func maxContiguousSum1(A []int) int {
	maxSum, n := 0, len(A)
	for i := 1; i < n; i++ { // for each possible start point
		for j := i; j < n; j++ { // for each possible end point
			currentSum := 0
			for k := i; k <= j; k++ {
				currentSum += A[k]
			}
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}
	return maxSum
}

// O(n^2)
func maxContiguousSum2(A []int) int {
	maxSum, n := 0, len(A)
	for i := 1; i < n; i++ {
		currentSum := 0
		for j := i; j < n; j++ {
			currentSum += A[j]
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}
	return maxSum
}

// O(n), DP
func maxContiguousSumDP(A []int) int {
	n := len(A)
	M := make([]int, n+1)
	maxSum := 0
	if A[0] > 0 {
		M[0] = A[0]
	} else {
		M[0] = 0
	}
	for i := 1; i < n; i++ {
		if M[i-1]+A[i] > 0 {
			M[i] = M[i-1] + A[i]
		} else {
			M[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		if M[i] > maxSum {
			maxSum = M[i]
		}
	}
	return maxSum
}

// O(n)
func maxContiguousSum4(A []int) int {
	sumSoFar, sumEndingHere, n := 0, 0, len(A)
	for i := 1; i < n; i++ {
		sumEndingHere = sumEndingHere + A[i]
		if sumEndingHere < 0 {
			sumEndingHere = 0
			continue
		}
		if sumSoFar < sumEndingHere {
			sumSoFar = sumEndingHere
		}
	}
	return sumSoFar
}

func main() {
	fmt.Println(maxContiguousSum1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxContiguousSum2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxContiguousSumDP([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxContiguousSum4([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
