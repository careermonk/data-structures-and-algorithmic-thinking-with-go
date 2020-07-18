# Copyright (c) July 12, 2020 CareerMonk Publications and others.
# E-Mail           		: info@careermonk.com 
# Creation Date    		: 2020-07-12 06:15:46 
# Last modification		: 2020-07-12 
#               by		: Narasimha Karumanchi 
# Book Title			: Data Structures And Algorithmic Thinking With Go
# Warranty         		: This software is provided "as is" without any 
# 				   warranty; without even the implied warranty of 
# 				    merchantability or fitness for a particular purpose.

package main

import (
	"fmt"
)

func findMin(A []int) int {
	// If the list has just one element then return that element.
	if len(A) == 1 {
		return A[0]
	}

	// initializing low and high pointers.
	low, high := 0, len(A)-1

	// if the last element is greater than the first element then there is no rotation.
	// e.g. 1 < 2 < 3 < 4 < 5 < 7. Already sorted array.
	// Hence the smallest element is first element. A[0]
	if A[high] > A[0] {
		return A[0]
	}

	// Binary search way
	for high >= low {
		// Find the mid element
		mid := low + (high-low)/2

		// if the mid element is greater than its next element then mid+1 element is the smallest
		// This point would be the point of change. From higher to lower value.
		if A[mid] > A[mid+1] {
			return A[mid+1]
		}

		// if the mid element is lesser than its previous element then mid element is the smallest
		if A[mid-1] > A[mid] {
			return A[mid]
		}

		// if the mid elements value is greater than the 0th element this means the least value
		// is still somewhere to the high as we are still dealing with elements greater than A[0]
		if A[mid] > A[0] {
			low = mid + 1
		} else {
			// if A[0] is greater than the mid value then this means the smallest value is somewhere to the low
			high = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}
