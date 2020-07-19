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

import "fmt"

type Matrix [][]int

func multiply(A, B Matrix) (C Matrix, ok bool) {
	rows, cols, extra := len(A), len(B[0]), len(B)
	if len(A[0]) != extra {
		return nil, false
	}
	C = make(Matrix, rows)
	for i := 0; i < rows; i++ {
		C[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			for k := 0; k < extra; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C, true
}

func (m Matrix) String() string {
	rows := len(m)
	cols := len(m[0])
	out := "["
	for r := 0; r < rows; r++ {
		if r > 0 {
			out += ",\n "
		}
		out += "[ "
		for c := 0; c < cols; c++ {
			if c > 0 {
				out += ", "
			}
			out += fmt.Sprintf("%4d", m[r][c])
		}
		out += " ]"
	}
	out += "]"
	return out
}

func main() {
	A := Matrix{[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8}}
	B := Matrix{[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
		[]int{10, 11, 12}}
	P, ok := multiply(A, B)
	if !ok {
		panic("Invalid dimensions")
	}
	fmt.Printf("Matrix A:\n%s\n\n", A)
	fmt.Printf("Matrix B:\n%s\n\n", B)
	fmt.Printf("Product of A and B:\n%s\n\n", P)
}
