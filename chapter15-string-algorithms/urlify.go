// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
//

package main

import (
	"fmt"
)

func urlify(A string) string {
	space, n := 0, len(A)
	for i := 0; i < n; i++ {
		if string(A[i]) == " " {
			space++
		}
	}
	originalLength := space*2 + n
	r := make([]byte, originalLength)
	for i := 0; i < n; i++ {
		r[i] = A[i]
	}
	fmt.Println(string(r))
	for i := n - 1; i >= 0; i-- {
		if r[i] == ' ' {
			r[originalLength-1] = '0'
			r[originalLength-2] = '2'
			r[originalLength-3] = '%'
			originalLength = originalLength - 3
		} else {
			r[originalLength-1] = r[i]
			originalLength--
		}
	}
	return string(r)
}

func main() {
	fmt.Println(urlify("CareerMonk Publications! "))
}
