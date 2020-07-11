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

// Interpolation Search in Golang
package main
import "fmt"
 
func interpolationSearch(A []int, data int) int {
	low := 0
	high := len(A) - 1
	for low <= high && data >= A[low] && data <= A[high] {
		guessIndex := low + (data-A[low])*(high-low)/(A[high]-A[low])
		if A[guessIndex] == data {
			return guessIndex
		} else if A[guessIndex] < data {
			low = guessIndex + 1
		} else if A[guessIndex] > data {
			high = guessIndex - 1
		}
	}
	return -1
}
 
func main(){
    items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
    fmt.Println(interpolationSearch(items,63))
}
