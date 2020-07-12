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

func mergeSort(A []int) []int {
	if len(A) <= 1 {
		return A
	}
	middle := len(A) / 2
	left := mergeSort(A[:middle])
	right := mergeSort(A[middle:])

	return merge(left, right)
}
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	for i := 0; len(left) > 0 || len(right) > 0; i++ {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}
	return result
}

func main() {
	A := []int{3, 4, 5, 2, 1}
	A = mergeSort(A)
	fmt.Println("\n After Insertion Sorting")
	for _, val := range A {
		fmt.Println(val)
	}
}
