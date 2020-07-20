package main

import (
	"fmt"
)

func maxSumWithNoTwoContinuousNumbers(A []int) int {
	n := len(A)
	M := make([]int, n+1)
	M[0] = A[0]
	if A[0] > A[1] {
		M[1] = A[0]
	} else {
		M[1] = A[1]
	}

	for i := 2; i < n; i++ {
		if M[i-1] > M[i-2]+A[i] {
			M[i] = M[i-1]
		} else {
			M[i] = M[i-2] + A[i]
		}
	}
	return M[n-1]
}

func main() {
	fmt.Println(maxSumWithNoTwoContinuousNumbers([]int{1, 2, 9, 4, 5, 0, 4, 11, 6}))
}
