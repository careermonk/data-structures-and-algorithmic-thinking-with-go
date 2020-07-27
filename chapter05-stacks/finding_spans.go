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

func findingSpans1(A []int) []int {
	n := len(A)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		j := 1
		for j <= i && A[i] > A[i-j] {
			j = j + 1
		}
		result[i] = j
	}
	return result
}

type Stack struct {
	top      int
	capacity uint
	array    []int
}

// Returns an initialized stack
func (stack *Stack) Init(capacity uint) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.array = make([]int, capacity)
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

	target := make([]int, stack.capacity)
	copy(target, stack.array[:stack.top+1])
	stack.array = target
}

func (stack *Stack) Push(data int) error {
	if stack.IsFull() {
		stack.Resize()
	}
	stack.top++
	stack.array[stack.top] = data
	return nil
}

func (stack *Stack) Pop() int {
	if stack.IsEmpty() {
		return -1
	}
	temp := stack.array[stack.top]
	stack.top--
	if stack.Size() < stack.capacity/2 {
		stack.Resize()
	}
	return temp
}

func (stack *Stack) Peek() int {
	if stack.IsEmpty() {
		return -1
	}
	temp := stack.array[stack.top]
	return temp
}

// Drain removes all elements that are currently in the stack.
func (stack *Stack) Drain() {
	stack.array = nil
	stack.top = -1
}

func findingSpans2(A []int) []int {
	n := len(A)
	result := make([]int, n)
	stack := NewStack(1)
	var P int

	for i := 0; i < n; i++ {
		for !stack.IsEmpty() && A[i] > A[stack.Peek()] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			P = -1
		} else {
			P = stack.Peek()
		}
		result[i] = i - P
		stack.Push(i)
	}
	return result
}

func main() {
	fmt.Println(findingSpans1([]int{6, 3, 4, 5, 2}))
	fmt.Println(findingSpans2([]int{6, 3, 4, 5, 2}))
}
