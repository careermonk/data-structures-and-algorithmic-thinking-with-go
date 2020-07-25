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
	"math"
	"sort"
)

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func twoElementsWithMinSumCloseToZero1(A []int) []int {
	n := len(A)
	if n < 2 {
		return []int{}
	}
	// Initialization of values
	min_i, min_j := 0, 1
	minSum := A[0] + A[1]
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum := A[i] + A[j]
			if abs(minSum) > abs(sum) {
				minSum = sum
				min_i = i
				min_j = j
			}
		}
	}
	return []int{A[min_i], A[min_j]}
}
func twoElementsWithMinSumCloseToZero2(A []int) []int {
	n := len(A)
	if n < 2 {
		return []int{}
	}

	// Variables to keep track of current sum and minimum sum
	min_sum := math.MaxInt32

	// left and right index variables
	i, j := 0, n-1

	// variable to keep track of the left and right pair for min_sum
	min_i, min_j := i, n-1

	/* Sort the elements */
	sort.Ints(A)

	for i < j {
		sum := A[i] + A[j]

		/*If abs(sum) is less then update the result items*/
		if abs(sum) < abs(min_sum) {
			min_sum = sum
			min_i = i
			min_j = j
		}
		if sum < 0 {
			i++
		} else {
			j--
		}
	}
	return []int{A[min_i], A[min_j]}
}

func main() {
	A := []int{1, 60, -10, 70, -80, 85}
	fmt.Println(twoElementsWithMinSumCloseToZero1(A))
	fmt.Println(twoElementsWithMinSumCloseToZero2(A))

}

[-80 85]
[-80 85]

Program exited.
