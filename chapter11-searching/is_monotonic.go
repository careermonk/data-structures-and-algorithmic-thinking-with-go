# Copyright (c) July 12, 2020 CareerMonk Publications and others.
# E-Mail           		: info@careermonk.com 
# Creation Date    		: 2020-07-12 06:15:46 
# Last modification		: 2020-07-12 
#               by		: Narasimha Karumanchi 
# Book Title			: Data Structures And Algorithmic Thinking With Go
# Warranty         		: This software is provided "as is" without any 
# 				   warranty; without even the implied warranty of 
# 				    merchantability or fitness for a particular purpose. 

package main

import (
	"fmt"
)

func isMonotonic(A []int) bool {
	if len(A) < 2 {
		return true
	}

	isIncreasing := 0

	for i := 1; i < len(A); i++ {
		if isIncreasing == 0 {
			if A[i-1] > A[i] {
				isIncreasing = -1
			} else if A[i-1] < A[i] {
				isIncreasing = 1
			}
		}

		if isIncreasing == 1 && A[i-1] > A[i] {
			return false
		}
		if isIncreasing == -1 && A[i-1] < A[i] {
			return false
		}
	}

	return true
}
