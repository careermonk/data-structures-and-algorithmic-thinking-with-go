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

func maxIndexDiff(A []int) int {
	maxDiff, n := -1, len(A)

	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if A[j] > A[i] && maxDiff < (j-i) {
				maxDiff = j - i
			}
		}
	}
	return maxDiff
}

func main() {
	fmt.Println((maxIndexDiff([]int{9, 2, 3, 4, 5, 6, 7, 8, 18, 0})))
}
