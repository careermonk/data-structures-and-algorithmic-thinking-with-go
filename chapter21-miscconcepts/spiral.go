/*
# Copyright (c) July 12, 2020 CareerMonk Publications and others.
# E-Mail           		: info@careermonk.com
# Creation Date    		: 2020-07-12 06:15:46
# Last modification		: 2020-07-12
#               by		: Narasimha Karumanchi
# Book Title			: Data Structures And Algorithmic Thinking With Go
# Warranty         		: This software is provided "as is" without any
# 				   warranty; without even the implied warranty of
# 				    merchantability or fitness for a particular purpose. */

// Draw Spiral in Golang
package main

import (
	"fmt"
)

func spiral(A [][]int) {
	rowStart, rowEnd, columnStart, columnEnd := 0, len(A)-1, 0, len(A[0])-1

	for {
		// columnEnd
		for i := columnStart; i <= columnEnd; i++ {
			fmt.Printf("%d ", A[rowStart][i])
		}
		rowStart++
		if rowStart > rowEnd {
			break
		}

		// down
		for i := rowStart; i <= rowEnd; i++ {
			fmt.Printf("%d ", A[i][columnEnd])
		}
		columnEnd--
		if columnStart > columnEnd {
			break
		}

		// columnStart
		for i := columnEnd; i >= columnStart; i-- {
			fmt.Printf("%d ", A[rowEnd][i])
		}
		rowEnd--
		if rowStart > rowEnd {
			break
		}

		// up
		for i := rowEnd; i >= rowStart; i-- {
			fmt.Printf("%d ", A[i][columnStart])
		}
		columnStart++
		if columnStart > columnEnd {
			break
		}
	}
}
func main() {
	A := [][]int{
		{0, 1, 2, 3},   /*  initializers for row indexed by 0 */
		{4, 5, 6, 7},   /*  initializers for row indexed by 1 */
		{8, 9, 10, 11}, /*  initializers for row indexed by 2 */
	}
	spiral(A)
}
