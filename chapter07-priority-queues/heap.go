  
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
	"math"
	"math/rand"
	"time"
)

// Item - defines the interface for an element to be held by a Heap instance
type Item interface {
	Less(item Item) bool
}

// Heap - binary heap with support for min heap operations
type Heap struct {
	size int
	data []Item
}

// New - returns a pointer to an empty min-heap
func New() *Heap {
	return &Heap{}
}

// Heapify - returns a pointer to a  min-heap composed of the elements of 'items'
func Heapify(items []Item) *Heap {
	h := New()

	n := len(items)

	h.data = make([]Item, n)
	copy(h.data, items)

	i := int(math.Floor(float64(n) / 2.0))

	for i > 0 {
		h.percolateDown(i)
		i--
	}
	return h
}

// Insert - inserts 'item' into the Heap, maintaining the min-heap invariant
func (h *Heap) Insert(item Item) {
	if h.size == 0 {
		h.data = make([]Item, 1)
		h.data[0] = item
	} else {
		h.data = append(h.data, item)
	}
	h.size++
	h.percolateUp()
}

// Extract - removes and returns the 'item' at the top of the heap, maintaining the min-heap invariant
func (h *Heap) Extract() (Item, error) {
	n := h.size
	if n == 0 {
		return nil, fmt.Errorf("Unable to extract from empty Heap")
	}

	m := h.data[0]
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]
	h.size--

	if h.size > 0 {
		h.percolateDown(0)
	} else {
		h.data = nil
	}

	return m, nil
}

func (h *Heap) percolateUp() {
	idx := h.size - 1

	if idx <= 0 {
		return
	}

	for {
		p := parent(idx)
		if p < 0 || h.data[p].Less(h.data[idx]) {
			break
		}
		swap(h, p, idx)
		idx = p
	}
}

func (h *Heap) percolateDown(i int) {
	p := i

	for {
		l := leftChild(p)
		r := rightChild(p)

		s := p

		if l < h.size && h.data[l].Less(h.data[s]) {
			s = l
		}

		if r < h.size && h.data[r].Less(h.data[s]) {
			s = r
		}

		if s == p {
			break
		}

		swap(h, p, s)
		p = s
	}
}

func getMaximum(h *Heap) int {
	if h.size == 0 {
		return -1
	}
	return h.data[0]
}

func swap(h *Heap, i int, j int) {
	tmp := h.data[i]
	h.data[i] = h.data[j]
	h.data[j] = tmp
}

func parent(i int) int {
	return int(math.Floor(float64(i-1) / 2.0))
}

func leftChild(p int) int {
	return (2 * p) + 1
}

func rightChild(p int) int {
	return (2 * p) + 2
}

// Test code

type Int int

func (a Int) Less(b Item) bool {
	val, ok := b.(Int)
	return ok && a <= val
}

func verifyHeap(h *Heap) bool {
	queue := make([]Int, 1)
	queue[0] = 0
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		l := leftChild(int(p))
		r := rightChild(int(p))

		if l < h.size {
			if !h.data[p].Less(h.data[l]) {
				return false
			}
			queue = append(queue, Int(l))
		}

		if r < h.size {
			if !h.data[p].Less(h.data[r]) {
				return false
			}
			queue = append(queue, Int(r))
		}
	}
	return true
}

func verifyStrictlyIncreasing(h *Heap) (bool, []Item) {
	prev, _ := h.Extract()
	order := []Item{prev}
	for h.size > 0 {
		curr, _ := h.Extract()
		order = append(order, curr)
		if curr.Less(prev) {
			return false, order
		}
		prev = curr
		order = append(order, prev)
	}
	return true, order
}

func randomPerm(n int) []Item {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	ints := r.Perm(n)
	items := make([]Item, n)

	for idx, item := range ints {
		items[idx] = Int(item)
	}
	return items
}

func main() {
	items := randomPerm(20)
	hp := New()
	for _, item := range items {
		fmt.Println("Inserting an element into Heap: ", hp.data)
		hp.Insert(item)
	}
	if !verifyHeap(hp) {
		fmt.Println("invalid Heap: ", hp.data)
		return
	}

	if ok, order := verifyStrictlyIncreasing(hp); !ok {
		fmt.Println("invalid Heap extraction order: ", order)
	}
}
