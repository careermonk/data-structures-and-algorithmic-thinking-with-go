// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else { // recursive case: multiply n by (n -1) factorial
		return n * factorial(n-1)
	}
}

var factorials []int
func factorialDP(n int) int {
	if len(factorials) == 0 {
		factorials = make([]int, n+1)
	}
	if n == 0 || n == 1 {
		return 1
	} else if factorials[n] != 0 { //Already calculated case
		return factorials[n]
	} else { // recursive case: multiply n by (n -1) factorial
		factorials[n] = n * factorialDP(n-1)
		return factorials[n]
	}
}

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorialDP(5))
	fmt.Println(factorial(5))
	fmt.Println(factorial(5))
}
