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

func removeChars(str, charsToBeRemoved string) string {
	mymap := map[byte]bool{}
	result := ""
	for i := 0; i < len(charsToBeRemoved); i++ {
		mymap[charsToBeRemoved[i]] = true
	}
	for i := 0; i < len(str); i++ {
		if mymap[str[i]] == false {
			result = result + string(str[i])
		}
	}
	return result
}

func main() {
	str := "abcedabcd"
	fmt.Println(removeChars(str, "e"))
	fmt.Println(removeChars(str, "a"))
}
