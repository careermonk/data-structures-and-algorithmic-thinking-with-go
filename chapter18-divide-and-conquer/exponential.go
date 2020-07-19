// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

package main

import (
	"fmt"
)

func exponential(k, n int64) int64{
	if k == 0 {
		return 1
	} else if n == 1 {
		return k
	} else {
		if n%2 == 1 {
			a := exponential(k, n-1)
			return a * k
		} else {
			a := exponential(k, n/2)
			return a * a
		}
	}
}

func main() {
	fmt.Println(exponential(2, 3))
}
