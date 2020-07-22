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

func minimumJumps(A []int) int {
	n, step, start, end := len(A), 0, 0, 0
	for end < n-1 {
		step++
		maxend := end + 1
		for i := start; i <= end; i++ {
			if i+A[i] >= n-1 {
				return step
			}
			maxend = max(maxend, i+A[i])
		}
		start = end + 1
		end = maxend
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	fmt.Println(minimumJumps([]int{2, 3, 1, 1, 4}))
}
