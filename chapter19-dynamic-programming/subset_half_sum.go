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

func subsetHalfSum(A []int) bool {
	K, n := 0, len(A)

	for i := 0; i < n; i++ {
		K += A[i]
	}
	T := make([]bool, K+1)
	T[0] = true // initialize the table
	for i := 1; i <= K; i++ {
		T[i] = false
	}
	// process the numbers one by one
	for i := 0; i < n; i++ {
		for j := K - A[i]; j >= 0; j-- {
			if T[j] {
				T[j+A[i]] = true
			}
		}
	}
	return T[K/2]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func subsetHalfSum2(A []int) bool {
	K, n := 0, len(A)

	for i := 0; i < n; i++ {
		K += A[i]
	}
	T := make([]bool, K+1)
	T[0] = true // initialize the table
	for i := 1; i <= K; i++ {
		T[i] = false
	}

	sort.Ints(A) // sort the array
	R := 0
	// process the numbers one by one
	for i := 0; i < n; i++ {
		for j := R; j >= 0; j-- {
			if T[j] {
				T[j+A[i]] = true
				R = min(K/2, R+A[i])
			}
		}
	}
	return T[K/2]
}

func main() {
	fmt.Println(subsetHalfSum([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(subsetHalfSum([]int{7, 2, 3, 4, 5, 6, 9}))
	fmt.Println(subsetHalfSum([]int{7, 2}))

	fmt.Println(subsetHalfSum2([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(subsetHalfSum2([]int{7, 2, 3, 4, 5, 6, 9}))
	fmt.Println(subsetHalfSum2([]int{7, 2}))
}
