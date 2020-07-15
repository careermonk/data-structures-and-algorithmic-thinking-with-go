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

func checkWhoWinsTheElection(A []int) int {
	counter := 0
	maxCounter := 0
	var candidate int
	candidate = A[0]
	for i := 0; i < len(A); i++ {
		candidate = A[i]
		counter = 0
		for j := i + 1; j < len(A); j++ {
			if A[i] == A[j] {
				counter++
			}
		}
		if counter > maxCounter {
			maxCounter = counter
			candidate = A[i]
		}
	}
	return candidate
}

// Driver code
func main() {
	votes := []int{2, 2, 4, 4, 7, 2, 7, 6, 4, 2}
	fmt.Print(checkWhoWinsTheElection(votes))
}
