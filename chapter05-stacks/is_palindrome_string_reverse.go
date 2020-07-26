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

func reverse(str string) string {
	// Create a slice of runes.
	result := []rune(str)

	// Beginning index of rune slice.
	var beg int

	// Ending index of the rune slice.
	end := len(result) - 1

	// Loop until bd and end meet in the middle of the slice.
	for beg < end {
		// Swap rune
		result[beg], result[end] = result[end], result[beg]
		beg = beg + 1
		end = end - 1
	}
	return string(result)
}

func isPalindrome(testString string) bool {
	if reverse(testString) == testString {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome("abcedabcd"))
	fmt.Println(isPalindrome("abccba"))
}
