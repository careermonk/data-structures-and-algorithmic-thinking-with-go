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

func lengthOfLIS(A []int) int {
	if len(A) == 0 {
		return 0
	}

	L := make([]int, len(A))
	L[0] = A[0]

	hi := 0

	for i := 1; i < len(A); i++ {
		idx := binarySearch(L, A[i], 0, hi)
		L[idx] = A[i]
		if idx == hi+1 {
			hi++
		}
	}

	return hi + 1
}

func binarySearch(L []int, data, lo, hi int) int {
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if L[mid] > data {
			hi = mid - 1
		} else if L[mid] < data {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
