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

func moveZerosToEnd(A []int) {
	i := 0
	j := 0
	for ; i < len(A); i++ {
		if A[i] != 0 {
			A[j], A[i] = A[i], A[j]
			j++
		}
	}
}

func main() {
	var A = []int{1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0}
	fmt.Println("Input is:  ", A)
	moveZerosToEnd(A)
	fmt.Println("Output is: ", A)
}
