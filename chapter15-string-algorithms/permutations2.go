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

// Permute the values at index i to len(str)-1
func permute(str []rune, i int) {
	if i == len(str) {
		fmt.Println(string(str))
	} else {
		for j := i; j < len(str); j++ {
			str[i], str[j] = str[j], str[i]
			permute(str, i+1)
			str[i], str[j] = str[j], str[i]
		}
	}
}

// Perm calls f with each permutation of str
func permutations(str string) {
	permute([]rune(str), 0)
}

func main() {
	permutations("abc")
}
