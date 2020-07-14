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

func moveSpacesToEnd(A []rune) {
	var n int = len(A) - int(1)
	count := 0
	for i := 0; i <= n; i++ {
		if A[i] != ' ' {
			// Count of non-space elements
			A[count] = A[i]
			count++
		}
	}
	for count <= n {
		A[count] = ' '
		count++
	}
}

func main() {
	var sparr []rune = []rune("move these spaces to end")
	fmt.Println("Input is:  ", string(sparr))
	moveSpacesToEnd(sparr)
	fmt.Println("Output is: ", string(sparr))
}
