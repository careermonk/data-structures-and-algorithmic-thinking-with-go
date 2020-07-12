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

package main
import "fmt"

func CheckDuplicatesInArray(A []int) bool {
	for i := 0; i < len(A); i++ {
		// Scan slice for a previous element of the same value.
		for v := 0; v < i; v++ {
			if A[v] == A[i] {
				return true
				break
			}
		}
	}
	return false
}

func main() {
	A := []int{10, 20, 30, 10, 10, 20, 40}
	fmt.Println(A)

	// Test the nested loop method.
	result := CheckDuplicatesInArray(A)
	fmt.Println(result)

	A = []int{10, 20, 30, 40, 50, 280, 940}
	fmt.Println(A)

	// Test the nested loop method.
	result = CheckDuplicatesInArray(A)
	fmt.Println(result)
}
