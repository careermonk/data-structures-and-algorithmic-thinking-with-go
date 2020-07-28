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
)

// Function KMP performing the Knuth-Morris-Pratt algorithm.
// Prints whether the P(pattern) was found and on what position in the text(T) or not.
// m - current match in T, i - current character in w, c - amount of comparisons.
func KMP(T, P string) {
	m, i, c := 0, 0, 0
	F := KMP_PrefixTable(P)
	for m+i < len(T) {
		fmt.Printf("\ncomparing characters %c %c at positions %d %d", T[m+i], P[i], m+i, i)
		c++
		if P[i] == T[m+i] {
			fmt.Printf(" - match")
			if i == len(P)-1 {
				fmt.Printf("\n\nWord %q was found at position %d in %q. \n%d comparisons were done.", P, m, T, c)
				return
			}
			i++
		} else {
			m = m + i - F[i]
			if F[i] > -1 {
				i = F[i]
			} else {
				i = 0
			}
		}
	}
	fmt.Printf("\n\nWord was not found.\n%d comparisons were done.", c)
	return
}

// Table building algorithm. Takes pattern P to be analyzed and return prefix table.
func KMP_PrefixTable(P string) (F []int) {
	F = make([]int, len(P))
	pos, cnd := 2, 0
	F[0], F[1] = -1, 0
	for pos < len(P) {
		if P[pos-1] == P[cnd] {
			cnd++
			F[pos] = cnd
			pos++
		} else if cnd > 0 {
			cnd = F[cnd]
		} else {
			F[pos] = 0
			pos++
		}
	}
	return F
}

func main() {
	KMP("bacbabababacaca", "ababaca")
}
