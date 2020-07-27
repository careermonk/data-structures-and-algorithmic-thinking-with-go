// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
//

package main

import (
	"container/heap"
	"fmt"
)

// min heap
type Heap struct {
	items []int
}

func (h *Heap) Len() int           { return len(h.items) }
func (h *Heap) Less(i, j int) bool { return h.items[i] < h.items[j] }
func (h *Heap) Swap(i, j int)      { h.items[i], h.items[j] = h.items[j], h.items[i] }
func (h *Heap) Push(v interface{}) { h.items = append(h.items, v.(int)) }
func (h *Heap) Pop() interface{} {
	n := len(h.items)
	x := h.items[n-1]
	h.items = h.items[:n-1]
	return x
}

func findKSmallest(S []int, k int) []int {
	h := &Heap{items: []int{}}
	heap.Init(h)
	result := []int{}
	for i := 0; i < len(S); i++ {
		heap.Push(h, S[i])
	}

	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).(int))
	}
	return result
}

func main() {
	fmt.Println(findKSmallest([]int{11, -4, 7, 8, -10}, 3))
}
