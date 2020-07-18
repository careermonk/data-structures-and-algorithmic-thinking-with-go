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

func singleNumber1(A []int) int {
	if len(A) == 1 {
		return A[0]
	}

LOOP:
	for i, num := range A {
		for j, num1 := range A {
			if i == j {
				continue
			}
			if num == num1 {
				continue LOOP
			}
		}
		return num
	}
	return 0
}

func singleNumber2(A []int) int {
	sort.Ints(A)
	for i := 0; i < len(A)-1; i += 2 {
		if A[i] != A[i+1] {
			return A[i]
		}
	}
	return A[len(A)-1]
}

func singleNumber3(nums []int) int {
	m := map[int]struct{}{}
	for _, num := range nums {
		if _, found := m[num]; found {
			delete(m, num)
		} else {
			m[num] = struct{}{}
		}
	}
	for k := range m {
		return k
	}
	return -1
}

func singleNumber4(A []int) int {
	result := 0
	for _, n := range A {
		result ^= n
	}
	return result
}

func main() {
	A := []int{4, 1, 2, 1, 2}
	fmt.Println(A)

	// Test the nested loop method.
	fmt.Println(singleNumber1(A))
	fmt.Println(singleNumber2(A))
	fmt.Println(singleNumber3(A))
	fmt.Println(singleNumber4(A))

	A = []int{2, 2, 1}
	fmt.Println(A)

	// Test the nested loop method.
	fmt.Println(singleNumber1(A))
	fmt.Println(singleNumber2(A))
	fmt.Println(singleNumber3(A))
	fmt.Println(singleNumber4(A))
}
