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

func ShellSort(A []int) {
	n := len(A)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && A[j] < A[j-h]; j -= h {
				A[j], A[j-h] = A[j-h], A[j]
			}
		}
		h /= 3
	}
}

func main() {
	A := []int{3, 4, 5, 2, 1}
	A = insertionSort(A)
	fmt.Println("\n After Insertion Sorting")
	for _, val := range A {
		fmt.Println(val)
	}
}
