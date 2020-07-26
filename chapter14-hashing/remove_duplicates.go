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
	"sort"
)

func removeDuplicates1(str string) string {
	result := ""
	for i := 0; i < len(str); i++ {
		// Scan slice for a previous element of the same value.
		found := false
		for v := 0; v < i; v++ {
			if str[v] == str[i] {
				found = true
				break
			}
		}
		// If no previous element found, append this one.
		if !found {
			result = result + string(str[i])
		}
	}
	return result
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
func removeDuplicates2(str string) string {
	if len(str) == 0 {
		return str
	}
	str = SortString(str)
	result := string(str[0])
	for i := 1; i < len(str); i++ {
		if str[i] != str[i-1] {
			result = result + string(str[i])
		}

	}
	return result
}

func removeDuplicates3(str string) string {
	mymap := map[byte]bool{}
	result := ""
	for i := 0; i < len(str); i++ {
		if mymap[str[i]] == true {
		} else {
			mymap[str[i]] = true
			result = result + string(str[i])
		}
	}
	return result
}

func main() {
	str := "abcdabcd"

	result := removeDuplicates1(str)
	fmt.Println(result)
	result = removeDuplicates2(str)
	fmt.Println(result)
	result = removeDuplicates3(str)
	fmt.Println(result)
}
