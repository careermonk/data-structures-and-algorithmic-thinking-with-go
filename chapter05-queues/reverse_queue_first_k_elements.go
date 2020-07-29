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
	"bytes"
	"fmt"
)

type Stack struct {
	top      int
	capacity uint
	array    []interface{}
}

// Returns an initialized stack
func (stack *Stack) Init(capacity uint) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.array = make([]interface{}, capacity)
	return stack
}

// Returns an new stack
func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity)
}

// Returns the size of the size
func (stack *Stack) Size() uint {
	return uint(stack.top + 1)
}

// Stack is full when top is equal to the last index
func (stack *Stack) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

// Stack is empty when top is equal to -1
func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

func (stack *Stack) Resize() {
	if stack.IsFull() {
		stack.capacity *= 2
	} else {
		stack.capacity /= 2
	}

	target := make([]interface{}, stack.capacity)
	copy(target, stack.array[:stack.top+1])
	stack.array = target
}

func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		stack.Resize()
	}
	stack.top++
	stack.array[stack.top] = data
	return nil
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	temp := stack.array[stack.top]
	stack.top--
	if stack.Size() < stack.capacity/2 {
		stack.Resize()
	}
	return temp
}

func (stack *Stack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	temp := stack.array[stack.top]
	return temp
}

// Drain removes all elements that are currently in the stack.
func (stack *Stack) Drain() {
	stack.array = nil
	stack.top = -1
}

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
func (q *Queue) Size() int {
	return q.size
}

// isEmpty returns true if the queue q has no elements.
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// isFull returns true if the queue q is at capacity.
func (q *Queue) IsFull() bool {
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

// Resize adjusts the size of queue q's underlying slice.
func (q *Queue) Resize() {
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
	q.rear = size - 1
}

// enQueue inserts a new value v at the rear of queue q.
func (q *Queue) EnQueue(v interface{}) {
	if q.IsFull() {
		q.Resize()
	}
	q.rear = (q.rear + 1) % q.capacity
	q.array[q.rear] = v
	if q.front == -1 {
		q.front = q.rear
	}
	q.size++
}

func reverseQueueFirstKElements(q *Queue, k int) {
	if q == nil || k > q.Size() {
		return
	} else if k > 0 {
		stack := NewStack(1)
		for i := 0; i < k; i++ {
			stack.Push(q.DeQueue())
		}
		for !stack.IsEmpty() {
			q.EnQueue(stack.Pop())
		}
		for i := 0; i < q.Size()-k; i++ { // wrap around rest of elements
			q.EnQueue(q.DeQueue())
		}
	}
}

func main() {
	var q Queue
	q.Init(1)
	for i := 11; i < 21; i++ {
		q.EnQueue(i)
	}
	reverseQueueFirstKElements(&q, 5)
	fmt.Println(q.String())
}
