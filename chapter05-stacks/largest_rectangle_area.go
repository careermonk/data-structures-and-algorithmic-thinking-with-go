// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.
package main

import (
	"fmt"
)

func findMin(A []int, i, j int) int {
	min := A[i]
	for i <= j {
		if min > A[i] {
			min = A[i]
		}
		i++
	}
	return min
}
func largestRectangleArea1(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for j, minimum_height := i, heights[i]; j < len(heights); j++ {
			minimum_height = findMin(heights, i, j)
			maxArea = findMax(maxArea, (j-i+1)*minimum_height)
		}
	}
	return maxArea
}

func largestRectangleArea2(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for j, minimum_height := i, heights[i]; j < len(heights); j++ {
			minimum_height = findMini(minimum_height, heights[j])
			maxArea = findMax(maxArea, (j-i+1)*minimum_height)
		}
	}
	return maxArea
}

func largestRectangleArea3(heights []int) int {
	i, max := 0, 0
	stack := make([]int, 0)

	for i < len(heights) {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[pop]
			var w int
			if len(stack) == 0 {
				w = i
			} else {
				w = i - stack[len(stack)-1] - 1
			}
			max = findMax(max, h*w)
		}
	}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		h := heights[pop]
		var w int
		if len(stack) == 0 {
			w = i
		} else {
			w = i - stack[len(stack)-1] - 1
		}
		max = findMax(max, h*w)
	}
	return max
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestRectangleArea1([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea2([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea3([]int{2, 1, 5, 6, 2, 3}))
}
