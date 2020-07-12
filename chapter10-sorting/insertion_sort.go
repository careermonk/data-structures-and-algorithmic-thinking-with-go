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

import "fmt"

func insertionSort(A []int) []int {
	n := len(A)
	// Run the outer loop n - 1 times, from index 1 to n-1, as first element is already sorted
	// At the end of ith iteration, we have sorted list [0, i]
	for i := 1; i <= n-1; i++ {
		// Pick ith element and keep swapping with i-1th element if A[i] < A[i-1]
		j := i
		for j > 0 {
			// If value at index j is smaller than the one at j-1, swap them
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
			}
			j -= 1
		}
	}
	return A
}

func main() {
	A := []int{3, 4, 5, 2, 1}
	A = insertionSort(A)
	fmt.Println("\n After Insertion Sorting")
	for _, val := range A {
		fmt.Println(val)
	}
}
