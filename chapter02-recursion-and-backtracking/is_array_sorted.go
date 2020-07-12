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

func isSorted(A []int) bool {
	n := len(A)
	if n == 1 {
		return true
	}
	if A[n-1] < A[n-2] {
		return false
	}
	return isSorted(A[:n-1])
}

func main() {
	A := []int{10, 20, 23, 23, 45, 78, 88}
	fmt.Println(isSorted(A))
	A = []int{10, 20, 3, 23, 45, 78, 88}
	fmt.Println(isSorted(A))
}
