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

func LIS(A []int) int {
	n := len(A)
	if n == 1 || n == 0 {
		return n
	}

	L := make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = 1
	}

	max := 1
	for i := n-1; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			if A[i] < A[j] && L[i] < L[j]+1 {
				L[i] = L[j] + 1
			}
		}
	}
	for i := 0; i < n; i++ {
		if max < L[i] {
			max = L[i]
		}
	}
	return max
}

func main() {
	fmt.Println(LIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
