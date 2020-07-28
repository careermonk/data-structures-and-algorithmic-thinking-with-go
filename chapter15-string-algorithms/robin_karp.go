// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
//

// Rabin-Karp string search algorithm in Golang
package main

import (
	"fmt"
)

const base = 16777619

func robin_karp(T string, P string) int {
	n, m := len(T), len(P)

	if n < m || len(P) == 0 {
		return 0
	}

	var mult uint32 = 1 // mult = base^(m-1)
	for i := 0; i < m-1; i++ {
		mult = (mult * base)
	}

	hp := hash(P)
	h := hash(T[:m])
	for i := 0; i < n-m+1; i++ {
		if h == hp {
			return i
		}
		if i > 0 {
			h = h - mult*uint32(T[i-1])
			h = h*base + uint32(T[i+m-1])
		}
	}
	return -1
}

func hash(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h = (h*base + uint32(s[i]))
	}
	return h
}

func main() {
	fmt.Println(robin_karp("AABAACAADAABAABA", "BAAB"))
	fmt.Println(robin_karp("AABAACAADAABAABA", "BAAB"))
}
