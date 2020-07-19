package main

import (
	"fmt"
)

func shuffleArray1(A []int, left, right int) {
	n := len(A) / 2
	for i, q, k := 0, 1, n; i < n; i, k, q = i+1, k+1, q+1 {
		for j := k; j > i+q; j-- {
			A[j-1], A[j] = A[j], A[j-1]
		}
	}
}

func shuffleArray2(A []int, left, right int) {
	c := left + (right-left)/2
	q := 1 + left + (c-left)/2
	if left == right { //base case when the array has only one element
		return
	}
	k, i := 1, q
	for i <= c {
		//swap elements around the center
		A[i], A[c+k] = A[c+k], A[i]
		i++
		k++
	}
	shuffleArray2(A, left, c)    //Recursively call the function on the left and right
	shuffleArray2(A, c+1, right) //Recursively call the function on the right
}

func main() {
	A := []int{1, 3, 5, 7, 2, 4, 6, 8}
	fmt.Println(A)
	shuffleArray1(A, 0, len(A)-1)
	fmt.Println(A)

	A = []int{1, 3, 5, 7, 2, 4, 6, 8}
	fmt.Println(A)
	shuffleArray2(A, 0, len(A)-1)
	fmt.Println(A)
}
