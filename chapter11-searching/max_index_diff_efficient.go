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

import "fmt"

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxIndexDiff(A []int) int {
	n := len(A)
	leftMins := make([]int, n)
	rightMaxs := make([]int, n)

	// Construct leftMins[] such that leftMins[i] stores the minimum value from (A[0], A[1], ... A[i])
	leftMins[0] = A[0]
	for i := 1; i < n; i++ {
		leftMins[i] = min(A[i], leftMins[i-1])
	}

	// Construct rightMaxs[] such that rightMaxs[j] stores the maximum value from (A[j], A[j+1], ..A[n-1])
	rightMaxs[n-1] = A[n-1]
	for j := n - 2; j >= 0; j-- {
		rightMaxs[j] = max(A[j], rightMaxs[j+1])
	}

	// Traverse both arrays from left to right to find optimum j - i. This process is similar to merge() of MergeSort.
	i, j, maxDiff := 0, 0, -1
	for j < n && i < n {
		if leftMins[i] < rightMaxs[j] {
			maxDiff = max(maxDiff, j-i)
			j = j + 1
		} else {
			i = i + 1
		}
	}
	return maxDiff
}

func main() {
	fmt.Println(maxIndexDiff([]int{9, 2, 3, 4, 5, 6, 7, 8, 18, 0}))
}
