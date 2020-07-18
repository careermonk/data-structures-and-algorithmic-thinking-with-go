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

func peakIndexInMountainArray(A []int) int {
	low, high := 0, len(A)-1
	for low < high {
		mid := low + (high-low)/2
		if A[mid] < A[mid+1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{3, 4, 5, 1, 2}))
}
