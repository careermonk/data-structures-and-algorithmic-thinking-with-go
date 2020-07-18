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

func mySwap(A []rune, i int, j int) {
	var temp rune = A[i]
	A[i] = A[j]
	A[j] = temp
}

func moveSpacesToBegin(A []rune) {
	var n int = len(A) - int(1)
	count := n
	var i int = n
	for ; i >= 0; i-- {
		if A[i] != ' ' {
			A[count] = A[i]
			count--
		}
	}
	for count >= 0 {
		A[count] = ' '
		count--
	}
}

func main() {
	var sparr = []rune("move these spaces to beginning\x00")
	fmt.Println("Value of A is: ", string(sparr))
	moveSpacesToBegin(sparr)
	fmt.Println("Value of A is: ", string(sparr))
}
