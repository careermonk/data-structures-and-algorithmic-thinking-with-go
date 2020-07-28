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

import "strings"

func reverseWords(sentence string) string {
    var words []string
    var start, end int
    
    for end < len(sentence) {
        for start < len(sentence) && sentence[start] == ' ' {
            start++
        }
        if start == len(sentence) {
            break
        }
        end = start+1
        for end < len(sentence) && sentence[end] != ' ' {
            end++
        }
        words = append(words, sentence[start:end])
        start = end+1
    }
    
    reverse(words)
    return strings.Join(words, " ")
}

func reverse(words []string) {
    for i, j := 0, len(words)-1; i < j; i, j=i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
}

func main() {
	fmt.Println(reverseWords("Hello, world"))
	fmt.Println(reverseWords("CareerMonk Publications"))
}
