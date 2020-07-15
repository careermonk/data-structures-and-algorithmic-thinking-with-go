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
	"sort"
)

func checkWhoWinsTheElection(A []int) int {
	currentCounter := 1
	maxCounter := 1
	var currentCandidate, maxCandidate int
	currentCandidate = A[0]
	maxCandidate = 0

	sort.Ints(A)
	for i := 1; i < len(A); i++ {
		if A[i] == currentCandidate {
			currentCounter++
		} else {
			currentCandidate = A[i]
			currentCounter = 1
		}
		if currentCounter > maxCounter {
			maxCandidate = currentCandidate
			maxCounter = currentCounter
		}

	}

	return maxCandidate
}

// Driver code
func main() {
	votes := []int{2, 2, 4, 4, 7, 2, 7, 6, 4, 2}
	fmt.Print(checkWhoWinsTheElection(votes))
}
