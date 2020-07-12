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

func quickSort(A []int) {
	recursionSort(A, 0, len(A)-1)
}

func recursionSort(A []int, left int, right int) {
	if left < right {
		pivot := partition(A, left, right)
		recursionSort(A, left, pivot-1)
		recursionSort(A, pivot+1, right)
	}
}

func partition(A []int, left int, right int) int {
	for left < right {
		for left < right && A[left] <= A[right] {
			right--
		}
		if left < right {
			A[left], A[right] = A[right], A[left]
			left++
		}

		for left < right && A[left] <= A[right] {
			left++
		}
		if left < right {
			A[left], A[right] = A[right], A[left]
			right--
		}
	}
	return left
}

func main() {
	A := []int{3, 4, 5, 2, 1}
	A = quickSort(A)
	fmt.Println("\n After quickSort")
	for _, val := range A {
		fmt.Println(val)
	}
}
