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
	"errors"
	"fmt"
)

func replaceWithNearestGreaterElement1(A []int) []int {
	n := len(A)
	for i := 0; i < n; i++ {
		nextNearestGreater := math.MinInt32
		for j := i + 1; j < n; j++ {
			if A[j] > nextNearestGreater {
				nextNearestGreater = A[j]
			}
		}
		A[i] = nextNearestGreater
	}
	return A
}

func replaceWithNearestGreaterElement2(A []int) []int {
	greatest := math.MinInt32
	for i := len(A) - 1; i >= 0; i-- {
		temp := greatest
		if A[i] > greatest {
			greatest = A[i]
		}
		A[i] = temp
	}
	return A
}

func main() {
	fmt.Println(replaceWithNearestGreaterElement1([]int{17, 18, 5, 4, 6, 1}))
	fmt.Println(replaceWithNearestGreaterElement2([]int{17, 18, 5, 4, 6, 1}))
}
