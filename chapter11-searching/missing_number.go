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

func missingNumber1(A []int) int {
	for v := 0; v <= len(A); v++ {
		found := false
		for i := 0; i < len(A); i++ {
			if v == A[i] {
				found = true
				break
			}
		}
		if found == false {
			return v
		}
	}
	return -1
}

func missingNumber2(A []int) int {
	if len(A) == 0 {
		return 0
	}
	sort.Ints(A)

	if A[len(A)-1] != len(A) {
		return len(A)
	}

	for i := 0; i < len(A); i++ {
		if A[i] != i {
			return i
		}
	}

	return -1
}

func missingNumber3(A []int) int {
	var HT = map[int]bool{}
	for _, num := range A {
		HT[num] = true
	}
	for v := 0; v <= len(A); v++ {
		if _, ok := HT[v]; ok == false {
			return v
		}
	}
	return -1
}

func missingNumber4(A []int) int {
	sum, numsSum := len(A), 0
	for i := range A {
		sum += i
		numsSum += A[i]
	}
	return sum - numsSum
}

func missingNumber5(A []int) int {
	n := len(A)
	sum := 0

	for _, num := range A {
		sum += num
	}

	return ((n) * (n + 1) / 2) - sum
}

func missingNumber6(A []int) int {
	missing := len(A)
	for i := 0; i < len(A); i++ {
		missing ^= i ^ A[i]
	}
	return missing
}

func main() {
	A := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(A)

	// Test the nested loop method.
	fmt.Println(missingNumber1(A))
	fmt.Println(missingNumber2(A))
	fmt.Println(missingNumber3(A))
	fmt.Println(missingNumber4(A))
	fmt.Println(missingNumber5(A))
	fmt.Println(missingNumber6(A))

	A = []int{3, 0, 1}
	fmt.Println(A)

	// Test the nested loop method.
	fmt.Println(missingNumber1(A))
	fmt.Println(missingNumber2(A))
	fmt.Println(missingNumber3(A))
	fmt.Println(missingNumber4(A))
	fmt.Println(missingNumber5(A))
	fmt.Println(missingNumber6(A))
}
