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

import "fmt"

func firstRepeatedChar(str string) byte {
	n := len(str)
	counters := [256]int{}

	for i := 0; i < 256; i++ {
		counters[i] = 0
	}

	for i := 0; i < n; i++ {
		if counters[str[i]] == 1 {
			return str[i]
		}
		counters[str[i]]++
	}
	return byte(0)
}

func main() {
	fmt.Printf("%c", firstRepeatedChar("abacd"))
}
