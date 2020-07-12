/*# Copyright (c) July 12, 2020 CareerMonk Publications and others.
# E-Mail           		: info@careermonk.com
# Creation Date    		: 2020-07-12 06:15:46
# Last modification		: 2020-07-12
#               by		: Narasimha Karumanchi
# Book Title			: Data Structures And Algorithmic Thinking With Go
# Warranty         		: This software is provided "as is" without any
# 				   warranty; without even the implied warranty of
# 				    merchantability or fitness for a particular purpose. */

package main

import (
	"fmt"
)

func insertionSort(A []int) {
	for i := 0; i < len(A); i++ {
		temp := A[i]
		j := i - 1
		for ; j >= 0 && A[j] > temp; j-- {
			A[j+1] = A[j]
		}
		A[j+1] = temp
	}
}

func BucketSort(A []int, bucketSize int) []int {
	var max, min int
	for _, n := range A {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	nBuckets := int(max-min)/bucketSize + 1
	buckets := make([][]int, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}

	for _, n := range A {
		idx := int(n-min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]int, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertionSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}

func main() {
	A := []int{3, 4, 5, 2, 1}
	A = BucketSort(A, 2)
	fmt.Println("\n After BucketSort")
	for _, val := range A {
		fmt.Println(val)
	}
}
