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

import "fmt"

func findMedianSortedArrays(A []int, B []int) float64 {
	if l := len(A) + len(B); l%2 == 0 {
		return (findKth(A, B, l/2-1) + findKth(A, B, l/2)) / 2.0
	} else {
		return findKth(A, B, l/2)
	}
}

func findKth(A []int, B []int, k int) float64 {
	for {
		l1, l2 := len(A), len(B)
		m1, m2 := l1/2, l2/2

		if l1 == 0 {
			return float64(B[k])
		} else if l2 == 0 {
			return float64(A[k])
		} else if k == 0 {
			if n1, n2 := A[0], B[0]; n1 <= n2 {
				return float64(n1)
			} else {
				return float64(n2)
			}
		}

		if k <= m1+m2 {
			if A[m1] <= B[m2] {
				B = B[:m2]
			} else {
				A = A[:m1]
			}
		} else {
			if A[m1] <= B[m2] {
				A = A[m1+1:]
				k -= m1 + 1
			} else {
				B = B[m2+1:]
				k -= m2 + 1
			}
		}
	}
}

// Driver program to test above functions
func main() {
	A := []int{1, 2}
	B := []int{3, 4}
	fmt.Println(findMedianSortedArrays(A, B))
}
