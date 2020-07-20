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

import (
	"fmt"
)

type DeQueue struct {
	indexes []int
}

func (d *DeQueue) push(i int) {
	d.indexes = append(d.indexes, i)
}

func (d *DeQueue) getFirst() int {
	return d.indexes[0]
}

func (d *DeQueue) popFirst() {
	d.indexes = d.indexes[1:]
}

func (d *DeQueue) getLast() int {
	return d.indexes[len(d.indexes)-1]
}

func (d *DeQueue) popLast() {
	d.indexes = d.indexes[:len(d.indexes)-1]
}

func (d *DeQueue) empty() bool {
	return 0 == len(d.indexes)
}

func maxSlidingWindow(A []int, k int) []int {
	if len(A) < k || 0 == k {
		return make([]int, 0)
	} else if 1 == k {
		return A
	}

	var (
		// len(A)-k+1 this is the number of windows.
		// If input has length 3 and k is 1, we'll have 3 - 1 + 1 = 3 windows.
		res = make([]int, len(A)-k+1)
		dq  = &DeQueue{}
	)

	for i := range A {
		// we pop the first element if it's outside of the current window.
		if false == dq.empty() && (i-k == dq.getFirst()) {
			dq.popFirst()
		}

		// we pop all the elements that are smaller than the current one.
		for false == dq.empty() && A[dq.getLast()] < A[i] {
			dq.popLast()
		}

		// we push the current one to the window.
		dq.push(i)

		// if we reached at least the first window end.
		if i >= k-1 {
			// we add the current result that is the first in the DeQueue.
			res[i-k+1] = A[dq.getFirst()]
		}
	}

	return res
}

func main() {
	fmt.Print(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

}
