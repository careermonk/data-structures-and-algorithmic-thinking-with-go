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

var fib []int

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if len(fib) == 0 {
		fib = make([]int, n+1)
	}
	if fib[n] != 0 {
		return fib[n]
	}
	fib[n] = fibonacci(n-1) + fibonacci(n-2)
	return fib[n]
}

func fibonacciR(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacciR(n-1) + fibonacciR(n-2)
}

func fibonacciDP(n int) int {
	fib = make([]int, n+1)
	fib[0], fib[1] = 0, 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

func fibonacciFinal(n int) int {
	a, b, sum := 0, 1, 0
	for i := 1; i < n; i++ {
		sum = a + b
		a = b
		b = sum
	}
	return sum
}

func main() {
	fmt.Println(fibonacciR(7))
	fmt.Println(fibonacci(7))
	fmt.Println(fibonacciDP(7))
	fmt.Println(fibonacciFinal(7))
}
