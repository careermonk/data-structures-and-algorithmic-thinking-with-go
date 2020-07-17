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

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	cache := make([]int, n)
	cache[0], cache[1] = 1, 2

	for i := 2; i < n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n-1]
}

// Driver program to test above functions
func main() {
	fmt.Println("Number of ways to climb = ", climbStairs(4))
}
