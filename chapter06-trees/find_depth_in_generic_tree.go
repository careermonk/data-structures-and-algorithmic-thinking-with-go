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

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

// function to find the height of tree
func findHeight(P []int, n int) int {
	var maxDepth int
	for i := 0; i < n; i++ {
		var j int = i
		var currentDepth int = 1
		for P[j] != -1 {
			currentDepth++
			j = P[j]
		}
		maxDepth = max(maxDepth, currentDepth)
	}
	return maxDepth
}

// Test code
func main() {
	var P []int = []int{-1, 0, 1, 6, 6, 0, 0, 2, 7}
	var n int = int(36 / 4)
	var height int = findHeight(P, n)

	fmt.Println("Height of the given generic tree is: ", height)
	return
}
