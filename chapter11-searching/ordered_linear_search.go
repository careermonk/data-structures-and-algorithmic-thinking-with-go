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

func  orderedLinearSearch(A []int, data int) int {
  for i, item:= range A {
    if(item == data) {
      return i+1
    } else if (A[i] > data) {
      return 0
    }
  }
  return 0
}

func main() {
  A := []int {67,68,16,8,5,86,29,21,50}
  a := orderedLinearSearch(A, 5)
  fmt.Printf("The Source Array : %v\n", A)
  fmt.Printf("The element %v is found at %v location", 5, a)
}
