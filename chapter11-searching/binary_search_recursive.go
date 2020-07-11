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
 
// Recursive Binary Search
func binarySearch(data int, A []int) bool {
	low := 0
	high := len(A) - 1
	if low <= high {
		mid := ((high + low) / 2)
		if A[mid] > data {
			return binarySearch(data, A[:mid])
		} else if A[mid] < data {
			return binarySearch(data, A[mid+1:])
		} else {
			return true
		}
	}
	return false
}

func main(){
    items := []int{1,2, 9, 10, 11, 45, 73, 80, 200} // should be a sorted array
    fmt.Println(binarySearch(63, items))
    fmt.Println(binarySearch(73, items))
}
