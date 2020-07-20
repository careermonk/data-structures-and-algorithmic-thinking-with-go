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
	h.size = len(items)

	i := int(n / 2)

	for i >= 0 {
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

func getMinimum(h *Heap) (Item, error) {
	if h.size == 0 {
		return nil, fmt.Errorf("Unable to get element from from empty Heap")
	}
	return h.data[0], nil
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

func kThSmallestElement(h *Heap, k int) (Item, error) {
	if k >= h.size {
		return nil, fmt.Errorf("Wrong position")
	}
	for i := 1; i < k; i++ {
		h.Extract()
	}
	return h.Extract()
}

func main() {
	items := randomPerm(30)
	hp := Heapify(items)
	for i := 0; i < hp.size; i++ {
		fmt.Print(hp.data[i].(Int), " ")
	}
	fmt.Print("\n")
	element, _ := kThSmallestElement(hp, 6)
	fmt.Print("Kth Smallest element: ", element)
	fmt.Print("\n")
	for i := 0; i < hp.size; i++ {
		fmt.Print(hp.data[i].(Int), " ")
	}
}
