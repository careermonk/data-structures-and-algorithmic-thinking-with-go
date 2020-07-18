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
	"sort"
)

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkDuplicatesInArray1(A []int) bool {
	for i := 0; i < len(A); i++ {
		for v := 0; v < i; v++ { // Scan slice for a previous element of the same value.
			if A[v] == A[i] {
				return true
			}
		}
	}
	return false
}

func checkDuplicatesInArray2(A []int) bool {
	sort.Ints(A)
	for i := 0; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			return true
		}
	}
	return false
}

func checkDuplicatesInArray3(A []int) bool {
	var HT = map[int]bool{}
	for _, num := range A {
		if _, ok := HT[num]; ok {
			return true
		}
		HT[num] = true
	}
	return false
}

func checkDuplicatesInArray4(A []int) bool {
	for i := 0; i < len(A); i++ {
		if A[abs(A[i])] < 0 {
			return true
		} else {
			A[abs(A[i])] = -A[abs(A[i])]
		}
	}
	return false
}

func main() {
	A := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(A)

	// Test the nested loop method.
	fmt.Println(checkDuplicatesInArray1(A))
	fmt.Println(checkDuplicatesInArray2(A))
	fmt.Println(checkDuplicatesInArray3(A))
	fmt.Println(checkDuplicatesInArray4(A))

	A = []int{0, 1, 2, 3, 4}
	fmt.Println(A)

	// Test the nested loop method.
	// Test the nested loop method.
	fmt.Println(checkDuplicatesInArray1(A))
	fmt.Println(checkDuplicatesInArray2(A))
	fmt.Println(checkDuplicatesInArray3(A))
	fmt.Println(checkDuplicatesInArray4(A))

	A = []int{3, 2, 1, 2, 2, 3}
	fmt.Println(A)

	// Test the nested loop method.
	// Test the nested loop method.
	fmt.Println(checkDuplicatesInArray1(A))
	fmt.Println(checkDuplicatesInArray2(A))
	fmt.Println(checkDuplicatesInArray3(A))
	fmt.Println(checkDuplicatesInArray4(A))
}
