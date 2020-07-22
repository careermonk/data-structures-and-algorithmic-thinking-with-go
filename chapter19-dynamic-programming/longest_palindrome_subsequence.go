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

func longestPalindromeSubseq(s string) int {
	L := make([][]int, len(s))
	for i := range L {
		L[i] = make([]int, len(s))
	}
	return helper(s, 0, len(s)-1, L)
}

func helper(s string, lBound, rBound int, L [][]int) int {
	if lBound > rBound {
		return 0
	}
	if L[lBound][rBound] != 0 {
		return L[lBound][rBound]
	}
	if lBound == rBound {
		L[lBound][rBound] = 1
		return 1
	}
	if s[lBound] == s[rBound] {
		L[lBound][rBound] = 2 + helper(s, lBound+1, rBound-1, L)
	} else {
		lMax := helper(s, lBound, rBound-1, L)
		rMax := helper(s, lBound+1, rBound, L)
		if lMax > rMax {
			L[lBound][rBound] = lMax
		} else {
			L[lBound][rBound] = rMax
		}
	}
	return L[lBound][rBound]
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}
