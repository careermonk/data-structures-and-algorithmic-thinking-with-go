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

func CatalanNumberR(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for i := 1; i <= n; i++ {
		count += CatalanNumberR(i-1) * CatalanNumberR(n-i)
	}
	return count
}

func CatalanNumber(n int) int {
	C := make([]int, n+1)
	C[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			C[i] += (C[j-1] * C[i-j])
		}
	}
	return C[n]
}

func main() {
	fmt.Println(CatalanNumberR(3))
	fmt.Println(CatalanNumber(3))
}
