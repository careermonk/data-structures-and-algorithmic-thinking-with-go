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

func firstUniqChar1(s string) rune {
	m := make(map[rune]uint, len(s)) //preallocate the map size
	for _, r := range s {
		m[r]++
	}

	for _, r := range s {
		if m[r] == 1 {
			return r
		}
	}
	return rune(0)
}

func firstUniqChar2(str string) rune {
	for i := 0; i < len(str); i++ {
		repeated := false
		for j := 0; j < len(str); j++ {
			if i != j && str[i] == str[j] {
				repeated = true
				break
			}
		}
		if !repeated { // Found the first non-repeated character
			return rune(str[i])
		}
	}
	return rune(0)
}

func main() {
	str := "abcedabcd"

	result := firstUniqChar1(str)
	fmt.Printf("%c\n", result)
	result = firstUniqChar2(str)
	fmt.Printf("%c\n", result)
}
