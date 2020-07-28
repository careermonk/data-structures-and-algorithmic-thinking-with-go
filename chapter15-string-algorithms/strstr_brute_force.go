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

func strStr1(T, P string) int {
	if len(P) == 0 {
		return 0
	}
	i, j := 0, 0
	for i < len(T) {
		if T[i] == P[0] {
			for j = 1; j < len(P); j++ {
				if i+j >= len(T) || T[i+j] != P[j] {
					break
				}
			}
			if j == len(P) {
				return i
			}
		}
		i++
	}
	return -1
}

func strStr2(T, P string) int {
	if len(P) == 0 {
		return 0
	}
	length := len(P)
	for i := 0; i < len(T)-len(P)+1; i++ {
		if T[i:i+length] == P {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr1("AABAACAADAABAABA", "BAAB"))
	fmt.Println(strStr2("AABAACAADAABAABA", "BAAB"))
}
