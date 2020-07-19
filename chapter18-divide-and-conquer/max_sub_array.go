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

import "fmt"
import "math"

func maxSubArray(A []int) int {
	length := len(A)
	low := 0
	high := length - 1
	_, _, sum := findMaxSubArray(A, low, high)
	return sum
}

func findMaxSubArray(A []int, low int, high int) (int, int, int) {
	if high == low {

		return low, high, A[low]

	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := findMaxSubArray(A, low, mid)
		rightLow, rightHigh, rightSum := findMaxSubArray(A, mid+1, high)
		crossLow, crossHigh, crossSum := maxCrossingSubArray(A, low, high, mid)

		if (leftSum >= rightSum) && (leftSum >= crossSum) {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}
	}
}

func maxCrossingSubArray(A []int, low, high, mid int) (int, int, int) {
	leftSum := math.MinInt32
	rightSum := math.MinInt32
	maxLeft := 0
	maxRight := 0

	sum := 0
	for i := mid; i >= low; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += A[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	return maxLeft, maxRight, (leftSum + rightSum)
}

func main() {
	A := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(A)

	fmt.Println(maxSubArray(A))

	A = []int{4, -1, 2, 1}
	fmt.Println(A)

	fmt.Println(maxSubArray(A))
}
