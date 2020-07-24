// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
//

package main

import (
	"fmt"
)

func checkPairwiseSorted(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	for i := 0; i < len(A)-1; i += 2 {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkPairwiseSorted([]int{2, 5, 3, 7, 9, 11}))
}
