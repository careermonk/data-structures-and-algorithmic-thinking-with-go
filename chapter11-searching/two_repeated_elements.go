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
	"math"
)

func repeatedElements(A []int) {
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] == A[j] {
				fmt.Println(A[i])
			}
		}
	}
}

func repeatedElementsHashing(A []int) {
	counts := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		counts[A[i]]++
		if counts[A[i]] == 2 {
			fmt.Println(A[i])
		}
	}
}

func repeatedElementsMath(A []int) {
	S, P, n := 0, 1, len(A)-2

	// Calculate Sum and Product of all elements in A[]
	for i := 0; i < len(A); i++ {
		S = S + A[i]
		P = P * A[i]
	}

	S = S - n*(n+1)/2
	P = P / factorial(n)

	D := int(math.Sqrt(float64(S*S - 4*P)))

	x := (D + S) / 2
	y := (S - D) / 2
	fmt.Println(x, y)
}

func factorial(n int) int {
	factVal := 1
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= i
		}
	}
	return factVal
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func repeatedElementsNegationTechnique(A []int) {
	n := len(A)
	for i := 0; i < n; i++ {
		if A[abs(A[i])] < 0 {
			fmt.Println(abs(A[i]))
		} else {
			A[abs(A[i])] = A[abs(A[i])] * -1
		}
	}
}

func repeatedElementsXOR(A []int) {
	XOR, X, Y, n := 0, 0, 0, len(A)

	for i := 0; i < n; i++ { // Compute XOR of all elements in A[]
		XOR ^= A[i]
	}

	for i := 1; i <= n; i++ { // Compute XOR of all elements {1, 2 ..n}
		XOR ^= i
	}

	rightMostSetBit := XOR & ^(XOR - 1) // Get the rightmost set bit in rightMostSetBit

	// Now divide elements in two sets by comparing rightmost set
	for i := 0; i < n; i++ {
		if (A[i] & rightMostSetBit) != 0 {
			X = X ^ A[i] // XOR of first set in A[]
		} else {
			Y = Y ^ A[i]
		}
	}

	for i := 1; i <= n; i++ {
		if (i & rightMostSetBit) != 0 {
			X = X ^ i // XOR of first set in A[] and {1, 2, ...n }
		} else {
			Y = Y ^ i
		}
	}
	fmt.Println(X, Y)
}

func main() {
	A := []int{4, 2, 4, 5, 2, 3, 1, 7, 6, 8}
	repeatedElements(A)
	repeatedElementsHashing(A)
	repeatedElementsMath(A)
	repeatedElementsNegationTechnique(A)
	repeatedElementsXOR(A)
}
