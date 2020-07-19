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

func Multiply(m1, m2 Matrix) (m3 Matrix, ok bool) {
	rows, cols, extra := len(m1), len(m2[0]), len(m2)
	if len(m1[0]) != extra {
		return nil, false
	}
	m3 = make(Matrix, rows)
	for i := 0; i < rows; i++ {
		m3[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			for k := 0; k < extra; k++ {
				m3[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return m3, true
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
	P, ok := Multiply(A, B)
	if !ok {
		panic("Invalid dimensions")
	}
	fmt.Printf("Matrix A:\n%s\n\n", A)
	fmt.Printf("Matrix B:\n%s\n\n", B)
	fmt.Printf("Product of A and B:\n%s\n\n", P)
}
