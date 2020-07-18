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
import (
	"fmt"
	"unicode"
)

func mySwap(A []rune, i int, j int) {
	var temp rune = A[i]
	A[i] = A[j]
	A[j] = temp
}

func moveSpacesToBegin(A []rune) {
	var i int = len(A) - int(1)
	var j int = i
	for ; j >= 0; j-- {
		if unicode.IsSpace(A[j]) != true {
			mySwap(A, func() int {
				defer func() {
					i--
				}()
				return i
			}(), j)
		}
	}
}

func main() {
	var sparr = []rune("move these spaces to end")
	fmt.Println("Input is: ", string(sparr))
	moveSpacesToBegin(sparr)
	fmt.Println("Output is: ", string(sparr))
}
