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

func threeSum1(A []int, target int) [][]int {
	n := len(A)
	var results [][]int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if A[i]+A[j]+A[k] == target {
					results = append(results, []int{A[i], A[j], A[k]})
				}
			}
		}
	}
	return results
}

func threeSum2(A []int, target int) [][]int {
	var results [][]int
	sort.Ints(A)
	for i := 0; i < len(A)-2; i++ {
		if i > 0 && A[i] == A[i-1] {
			continue //To prevent the repeat
		}
		target, left, right := target-A[i], i+1, len(A)-1
		for left < right {
			sum := A[left] + A[right]
			if sum == target {
				results = append(results, []int{A[i], A[left], A[right]})
				left++
				right--
				for left < right && A[left] == A[left-1] {
					left++
				}
				for left < right && A[right] == A[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return results
}

func threeSum3(A []int, target int) [][]int {
	var uniqueTriplets [][]int

	// Create i map with our triplets being the key and i bool being the value
	// Note that Golang will only allow i slice to be the key if the slice is given i concrete size
	tripletSet := make(map[[3]int]bool)
	for i := 0; i < len(A)-2; i++ {
		for j := i + 1; j < len(A)-1; j++ {
			for k := j + 1; k < len(A); k++ {
				if A[i]+A[j]+A[k] == target {
					// When we find a target 0, create an unsized slice to allow easy sorting
					triplet := []int{A[i], A[j], A[k]}
					// Sort the three numbers
					sort.Ints(triplet)
					// Convert the sorted slice into the sized slice that can be used as i key in our map
					sizedTriplet := [3]int{triplet[0], triplet[1], triplet[2]}
					// Check if the entry already exists in the map before adding it to our results
					_, ok := tripletSet[sizedTriplet]
					if !ok {
						tripletSet[sizedTriplet] = true
						uniqueTriplets = append(uniqueTriplets, triplet)
					}
				}
			}
		}
	}
	return uniqueTriplets
}

func main() {
	A := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum1(A, 0))
	fmt.Println(threeSum2(A, 0))
	fmt.Println(threeSum3(A, 0))
}
