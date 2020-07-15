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

// Function to generate all k-ary strings
func generateK_aryStrings(n int, A []int, i int, k int) {
	if i == n {
		printResult(A, n)
		return
	}
	for j := 0; j < k; j++ {
		// assign j at ith position and try for all other permutations for remaining positions
		A[i] = j
		generateK_aryStrings(n, A, i+1, k)
	}
}

func main() {
	var n int = 4
	A := make([]int, n)

	// Print all binary strings
	generateK_aryStrings(n, A, 0, 3)
	return
}
