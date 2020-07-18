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
)

func majorityElement(A []int) int {
	var count, value int
	for i := range A {
		if count == 0 {
			value = A[i]
		}
		if value == A[i] {
			count++
		} else {
			count--
		}
	}
	return value
}
func main() {
	A := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(A)

	// Test the nested loop method.
	fmt.Println(majorityElement(A))
}
