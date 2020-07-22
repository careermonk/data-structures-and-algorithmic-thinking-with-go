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

// A utility function that returns true if there is a subset of A[] with sum equal to given sum
func subsetSum(A []int, n, sum int) bool {
	if sum == 0 {
		return true
	}
	if n == 0 && sum != 0 {
		return false
	}
	// If last element is greater than sum, then ignore it
	if A[n-1] > sum {
		return subsetSum(A, n-1, sum)
	}
	return subsetSum(A, n-1, sum) || subsetSum(A, n-1, sum-A[n-1])
}

// Returns true if A[] can be partitioned in two subsets of equal sum, otherwise false
func findPartition(A []int) bool {
	n := len(A)
	// Calculate sum of all elements
	sum := 0
	for i := 0; i < n; i++ {
		sum += A[i]
	}

	// If sum is odd, there cannot be two subsets with equal sum
	if sum%2 != 0 {
		return false
	}

	// Find if there is a subset with a sum equal to half of total sum
	return subsetSum(A, n, sum/2)
}

// Returns true if A[] can be partitioned in two subsets of equal sum, otherwise false
func findPartitionDP(A []int) bool {
	sum, n := 0, len(A)

	// calculate sum of all elements
	for i := 0; i < n; i++ {
		sum += A[i]
	}

	// If sum is odd, there cannot be two subsets with equal sum
	if sum%2 != 0 {
		return false
	}

	part := make([][]bool, sum/2+1)
	for i := range part {
		part[i] = make([]bool, n+1) // defaults to 0
	}

	// initialize top row as true
	for i := 0; i <= n; i++ {
		part[0][i] = true
	}

	// initialize leftmost column, except part[0][0], as 0
	for i := 1; i <= sum/2; i++ {
		part[i][0] = false
	}

	// Fill the partition table in bottom up manner
	for i := 1; i <= sum/2; i++ {
		for j := 1; j <= n; j++ {
			part[i][j] = part[i][j-1]
			if i >= A[j-1] {
				part[i][j] = part[i][j] || part[i-A[j-1]][j-1]
			}
		}
	}
	return part[sum/2][n]
}

func main() {
	fmt.Println(findPartition([]int{1, 5, 11, 5}))
	fmt.Println(findPartition([]int{1, 2, 3, 5}))

	fmt.Println(findPartitionDP([]int{1, 5, 11, 5}))
	fmt.Println(findPartitionDP([]int{1, 2, 3, 5}))
}
