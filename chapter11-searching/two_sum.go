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
	"sort"
)

func twoSum1(A []int, target int) []int {
	for i, v := range A {
		for j := i + 1; j < len(A); j++ {
			if A[j] == target-v {
				return []int{i, j}
			}
		}
	}
	panic("should never happen")
}

func twoSum2(A []int, target int) []int {
	sort.Ints(A)
	fmt.Println(A)
	i, j := 0, len(A)-1

	for i < j {
		sum := A[i] + A[j]
		if sum == target {
			break
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{i, j}
}

func twoSum3(A []int, K int) []int {
	H := make(map[int]int)
	for i, v := range A {
		k := K - v
		if _, ok := H[k]; ok {
			return []int{H[k], i}
		}
		H[v] = i
	}
	return nil
}

// Driver program to test above functions
func main() {
	fmt.Println(twoSum1([]int{7, 2, 11, 15}, 9))
	fmt.Println(twoSum2([]int{7, 2, 11, 1}, 9))
	fmt.Println(twoSum2([]int{7, 2, 11, 1}, 12))
	fmt.Println(twoSum3([]int{7, 2, 11, 1}, 9))
}
