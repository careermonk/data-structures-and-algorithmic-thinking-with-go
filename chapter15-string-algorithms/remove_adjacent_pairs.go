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

func removeAdjacentPairs(str string) string {
	r := []rune(str)
	n, i, j := len(r), 0,0
	for i = 0; i < n; i++ {
		if j > 0 && i < n && (r[i] == r[j-1]) { //Cancel pairs
			j--
		} else {
			r[j] = r[i]
			j++
		}
	}
	return string(r[:j])
}

func main() {
	fmt.Println(removeAdjacentPairs("abbaca"))
}
