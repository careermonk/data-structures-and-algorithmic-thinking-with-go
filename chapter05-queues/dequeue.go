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
	"fmt"
)

const minCapacity = 16

type DeQueue struct {
	buffer   []interface{}
	head     int
	tail     int
	size     int
	capacity int
}

func (q *DeQueue) Size() int {
	return q.size
}

func (q *DeQueue) IsFull() bool {
	return q.size == q.capacity
}

func (q *DeQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *DeQueue) PushBack(elem interface{}) {
	q.growIfFull()

	q.buffer[q.tail] = elem
	// Calculate new tail position.
	q.tail = q.next(q.tail)
	q.size++
}

func (q *DeQueue) PushFront(elem interface{}) {
	q.growIfFull()

	// Calculate new head position.
	q.head = q.prev(q.head)
	q.buffer[q.head] = elem
	q.size++
}

func (q *DeQueue) PopFront() interface{} {
	if q.size <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	ret := q.buffer[q.head]
	q.buffer[q.head] = nil
	// Calculate new head position.
	q.head = q.next(q.head)
	q.size--

	q.shrinkIfExcess()
	return ret
}

func (q *DeQueue) PopBack() interface{} {
	if q.size <= 0 {
		panic("deque: PopBack() called on empty queue")
	}

	// Calculate new tail position
	q.tail = q.prev(q.tail)

	// Remove value at tail.
	ret := q.buffer[q.tail]
	q.buffer[q.tail] = nil
	q.size--

	q.shrinkIfExcess()
	return ret
}

func (q *DeQueue) Front() interface{} {
	if q.size <= 0 {
		panic("deque: Front() called when empty")
	}
	return q.buffer[q.head]
}

func (q *DeQueue) Back() interface{} {
	if q.size <= 0 {
		panic("deque: Back() called when empty")
	}
	return q.buffer[q.prev(q.tail)]
}

func (q *DeQueue) At(i int) interface{} {
	if i < 0 || i >= q.size {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return q.buffer[(q.head+i)&(len(q.buffer)-1)]
}

func (q *DeQueue) Set(i int, elem interface{}) {
	if i < 0 || i >= q.size {
		panic("deque: Set() called with index out of range")
	}
	// bitwise modulus
	q.buffer[(q.head+i)&(len(q.buffer)-1)] = elem
}

func (q *DeQueue) Clear() {
	// bitwise modulus
	modBits := len(q.buffer) - 1
	for h := q.head; h != q.tail; h = (h + 1) & modBits {
		q.buffer[h] = nil
	}
	q.head = 0
	q.tail = 0
	q.size = 0
}

func (q *DeQueue) SetMinCapacity(minCapacityExp uint) {
	if 1<<minCapacityExp > minCapacity {
		q.capacity = 1 << minCapacityExp
	} else {
		q.capacity = minCapacity
	}
}

func (q *DeQueue) prev(i int) int {
	return (i - 1) & (len(q.buffer) - 1) // bitwise modulus
}

func (q *DeQueue) next(i int) int {
	return (i + 1) & (len(q.buffer) - 1) // bitwise modulus
}

func (q *DeQueue) growIfFull() {
	if len(q.buffer) == 0 {
		if q.capacity == 0 {
			q.capacity = minCapacity
		}
		q.buffer = make([]interface{}, q.capacity)
		return
	}
	if q.size == len(q.buffer) {
		q.resize()
	}
}

func (q *DeQueue) shrinkIfExcess() {
	if len(q.buffer) > q.capacity && (q.size<<2) == len(q.buffer) {
		q.resize()
	}
}

func (q *DeQueue) resize() {
	newBuf := make([]interface{}, q.size<<1)
	if q.tail > q.head {
		copy(newBuf, q.buffer[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buffer[q.head:])
		copy(newBuf[n:], q.buffer[:q.tail])
	}

	q.head = 0
	q.tail = q.size
	q.buffer = newBuf
}

func main() {
	var q DeQueue
	q.PushBack("foo")
	q.PushBack("bar")
	q.PushBack("baz")
	if q.Front() != "foo" {
		fmt.Println("wrong value at front of queue")
	}
	if q.Back() != "baz" {
		fmt.Println("wrong value at back of queue")
	}

	if q.PopFront() != "foo" {
		fmt.Println("wrong value removed from front of queue")
	}
	if q.Front() != "bar" {
		fmt.Println("wrong value remaining at front of queue")
	}
	if q.Back() != "baz" {
		fmt.Println("wrong value remaining at back of queue")
	}

	if q.PopBack() != "baz" {
		fmt.Println("wrong value removed from back of queue")
	}
	if q.Front() != "bar" {
		fmt.Println("wrong value remaining at front of queue")
	}
	if q.Back() != "bar" {
		fmt.Println("wrong value remaining at back of queue")
	}
}
