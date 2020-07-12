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

func findLargestElement(A []int) int {
	largestElement := 0
	for i := 0; i < len(A); i++ {
		if A[i] > largestElement {
			largestElement = A[i]
		}
	}
	return largestElement
}

func RadixSort(A []int) {
	largestElement := findLargestElement(A)
	size := len(A)
	significantDigit := 1
	semiSorted := make([]int, size, size)
	for largestElement/significantDigit > 0 {
		bucket := [10]int{0}
		for i := 0; i < size; i++ {
			bucket[(A[i]/significantDigit)%10]++
		}
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}
		for i := size - 1; i >= 0; i-- {
			bucket[(A[i]/significantDigit)%10]--
			semiSorted[bucket[(A[i]/significantDigit)%10]] = A[i]
		}
		for i := 0; i < size; i++ {
			A[i] = semiSorted[i]
		}
		significantDigit *= 10
	}
}

func main() {
	A := []int{123, 344, 589, 24, 167}
	RadixSort(A)
	fmt.Println("\n After RadixSort")
	for _, val := range A {
		fmt.Println(val)
	}
}
