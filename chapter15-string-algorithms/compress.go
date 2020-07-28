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
	"strconv"
)

func compress(S string) (int, string) {
	n := len(S)
	A := []rune(S)
	anchor := 0 // slow pointer
	read := 0   // fast pointer
	write := 0  // writing head pointer

	for read <= n {
		for (read < n) && (A[anchor] == A[read]) {
			read++
		}
		A[write] = A[anchor]
		write++
		count := read - anchor
		if count > 1 {
			s := strconv.Itoa(count)
			for k := 0; k < len(s); k++ {
				A[write] = rune(s[k])
				write++
			}
		}
		anchor = read
		read++
	}
	return write, string(A)
}

func main() {
	fmt.Println(compress("aabbccc"))
}
