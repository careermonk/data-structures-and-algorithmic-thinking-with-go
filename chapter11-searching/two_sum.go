// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

package main

import "fmt"

func twoSum(A []int, K int) []int {
	H := make(map[int]int)
	for i, v := range A {
		k := K - v
		if _, ok := H[k]; ok {
			return []int{H[k], i}
		}
		H[v] = i
	}
	return nil
}

// Driver program to test above functions
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
