/*
# Copyright (c) July 12, 2020 CareerMonk Publications and others.
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

func bubbleSort(A []int) []int {
	var sorted bool
	items := len(A)
	// We run the outer loop until we have sorted all the items.
	for !sorted {
		// In each iteration we are going to change sorted to true.
		sorted = true
		// Now we're going to range over our slice checking if they're sorted or not.
		for i := 1; i < items; i++ {
			// If they're not sorted we swap them and change sorted to false to loop again.
			if A[i-1] > A[i] {
				A[i-1], A[i] = A[i], A[i-1]
				sorted = false
			}
		}
	}
	return A
}

func main() {
	A := []int{3, 4, 5, 2, 1}
	A = bubbleSort(A)
	fmt.Println("\n After Bubble Sorting")
	for _, val := range A {
		fmt.Println(val)
	}
}
