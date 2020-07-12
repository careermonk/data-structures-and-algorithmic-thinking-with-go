/*# Copyright (c) July 12, 2020 CareerMonk Publications and others.
# E-Mail           		: info@careermonk.com
# Creation Date    		: 2020-07-12 06:15:46
# Last modification		: 2020-07-12
#               by		: Narasimha Karumanchi
# Book Title			: Data Structures And Algorithmic Thinking With Go
# Warranty         		: This software is provided "as is" without any
# 				   warranty; without even the implied warranty of
# 				    merchantability or fitness for a particular purpose. */

package main

import (
	"fmt"
)

func CountingSort(A []int, K int) []int {
	bucketLen := K + 1
	C := make([]int, bucketLen)

	sortedIndex := 0
	length := len(A)

	for i := 0; i < length; i++ {
		C[A[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for C[j] > 0 {
			A[sortedIndex] = j
			sortedIndex += 1
			C[j] -= 1
		}
	}
	return A
}

func main() {
	A := []int{3, 4, 5, 2, 1}
	A = CountingSort(A, 5)
	fmt.Println("\n After CountingSort")
	for _, val := range A {
		fmt.Println(val)
	}
}
