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

func printResult(A []int, n int) {
	var i int
	for ; i < n; i++ {
		// Function to print the output
		fmt.Print(A[i])
	}
	fmt.Printf("\n")
}

// Function to generate all binary strings
func generateBinaryStrings(n int, A []int, i int) {
	if i == n {
		printResult(A, n)
		return
	}
	// assign "0" at ith position and try for all other permutations for remaining positions
	A[i] = 0
	generateBinaryStrings(n, A, i+1)
	// assign "1" at ith position and try for all other permutations for remaining positions
	A[i] = 1
	generateBinaryStrings(n, A, i+1)
}

func main() {
	var n int = 4
	A := make([]int, n)

	// Print all binary strings
	generateBinaryStrings(n, A, 0)
	return
}
