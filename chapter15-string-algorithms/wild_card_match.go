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

func isMatch(text, pattern string) bool {
	prev, now := make([]bool, len(pattern)+1), make([]bool, len(pattern)+1)
	for i := 0; i <= len(text); i++ {
		now, prev = prev, now
		now[0] = i == 0
		for j := 1; j <= len(pattern); j++ {
			if pattern[j-1] == '*' {
				now[j] = prev[j] || prev[j-1] || now[j-1]
			} else {
				now[j] = prev[j-1] && (text[i-1] == pattern[j-1] || pattern[j-1] == '?')
			}
		}
	}
	return now[len(pattern)]
}

func main() {
	fmt.Println(isMatch("CareerMonk Publications", "*ca?ions"))
}
