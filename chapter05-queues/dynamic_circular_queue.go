/*# Copyright (c) July 12, 2020 CareerMonk Publications and others.
# E-Mail           		: info@careermonk.com
# Creation Date    		: 2020-07-12 06:15:46
# Last modification		: 2020-07-12
#               by		: Narasimha Karumanchi
# Book Title			: Data Structures And Algorithmic Thinking With Go
# Warranty         		: This software is provided "as is" without any
# 				   warranty; without even the implied warranty of
# 				    merchantability or fitness for a particular purpose. */

package main

import (
	"bytes"
	"fmt"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

type Queue struct {
	// enQueue increments rear then writes to array[rear] ; enQueue
	// writes to array[front] then decrements front; len(array) is a power
	// of two; unused slots are nil and not garbage.
	array    []interface{}
	front    int
	rear     int
	capacity int
	size     int
}

// New returns an initialized isEmpty queue.
func New(capacity int) *Queue {
	return new(Queue).Init(capacity)
}

// Init initializes with capacity
func (q *Queue) Init(capacity int) *Queue {
	q.array = make([]interface{}, capacity)
	q.front, q.rear, q.size, q.capacity = -1, -1, 0, capacity
	return q
}

// size returns the number of elements of queue q.
func (q *Queue) length() int {
	return q.size
}

// isEmpty returns true if the queue q has no elements.
func (q *Queue) isEmpty() bool {
	return q.size == 0
}

// isFull returns true if the queue q is at capacity.
func (q *Queue) isFull() bool {
	return q.size == q.capacity
}

// String returns a string representation of queue q formatted
// from front to rear.
func (q *Queue) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	j := q.front
	for i := 0; i < q.size; i++ {
		result.WriteString(fmt.Sprintf("%v", q.array[j]))
		if i < q.size-1 {
			result.WriteByte(' ')
		}
		j = (j + 1) % q.capacity
	}
	result.WriteByte(']')
	return result.String()
}

// Front returns the first element of queue q or nil.
func (q *Queue) Front() interface{} {
	// no need to check q.isEmpty(), unused slots are nil
	return q.array[q.front]
}

// Back returns the last element of queue q or nil.
func (q *Queue) Back() interface{} {
	return q.array[q.rear]
}

// resize adjusts the size of queue q's underlying slice.
func (q *Queue) resize() {
	size := q.capacity
	q.capacity = q.capacity * 2
	adjusted := make([]interface{}, q.capacity)
	if q.front < q.rear {
		// array not "wrapped" around, one copy suffices
		copy(adjusted, q.array[q.front:q.rear+1])
	} else {
		// array is "wrapped" around, need two copies
		n := copy(adjusted, q.array[q.front:])
		copy(adjusted[n:], q.array[:q.rear+1])
	}
	q.array = adjusted
	q.front = 0
	q.rear = size-1
}

// enQueue inserts a new value v at the rear of queue q.
func (q *Queue) enQueue(v interface{}) {
	if q.isFull() {
		q.resize()
	}
	q.rear = (q.rear + 1) % q.capacity
	q.array[q.rear] = v
	if q.front == -1 {
		q.front = q.rear
	}
	q.size++
}

// deQueue removes and returns the first element of queue q or MinInt.
func (q *Queue) deQueue() interface{} {
	if q.isEmpty() {
		return MinInt
	}
	data := q.array[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
		q.size = 0
	} else {
		q.front = (q.front + 1) % q.capacity
		q.size -= 1
	}
	return data
}

func main() {
	var q Queue
	q.Init(1)
	fmt.Println(q.String())
	q.enQueue(1)
	fmt.Println(q.String())
	q.enQueue(2)
	fmt.Println(q.String())
	q.enQueue(3)
	fmt.Println(q.String())
	q.enQueue(4)
	fmt.Println(q.String())
	q.enQueue(5)
	fmt.Println(q.String())
	q.enQueue(6)
	fmt.Println(q.String())
	q.enQueue(7)
	fmt.Println(q.String())
	q.enQueue(8)
	fmt.Println(q.String())
	q.enQueue(9)
	fmt.Println(q.String())
	q.enQueue(10)
	fmt.Println(q.String())
	q.enQueue(11)
	fmt.Println(q.String())
	q.enQueue(12)
	fmt.Println(q.String())
	q.enQueue(13)
	fmt.Println(q.String())
	want := "[1 2 3 4 5 6 7 8 9 10 11 12 13]"
	if s := q.String(); s != want {
		fmt.Println("q.String() = ", s, "want = ", want)
	}

	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	want = "[]"
	fmt.Println(q.String())
	if s := q.String(); s != want {
		fmt.Println("q.String() = ", s, "want = ", want)
	}
}
