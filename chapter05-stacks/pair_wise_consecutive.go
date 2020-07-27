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
	"math"
)

func replaceWithNearestGreaterElement1(A []int) []int {
	n := len(A)
	for i := 0; i < n; i++ {
		nextNearestGreater := math.MinInt32
		for j := i + 1; j < n; j++ {
			if A[j] > nextNearestGreater {
				nextNearestGreater = A[j]
			}
		}
		A[i] = nextNearestGreater
	}
	return A
}

func replaceWithNearestGreaterElement2(A []int) []int {
	greatest := math.MinInt32
	for i := len(A) - 1; i >= 0; i-- {
		temp := greatest
		if A[i] > greatest {
			greatest = A[i]
		}
		A[i] = temp
	}
	return A
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func pairWiseConsecutive(stack *Stack) bool {
	// Transfer elements of s to aux.
	auxStack := NewStack(1)
	for !stack.IsEmpty() {
		auxStack.Push(stack.Peek())
		stack.Pop()
	}

	// Traverse auxauxStack and see if elements are pairwise consecutive
	// or not. We also need to make sure that original content is retained.
	result := true
	for auxStack.Size() > 1 {
		// Fetch current top two elements of aux and check if they are consecutive.
		x := auxStack.Peek()
		auxStack.Pop()
		y := auxStack.Peek()
		auxStack.Pop()
		if abs(x-y) != 1 {
			result = false
		}

		// Push the elements to original stack
		stack.Push(x)
		stack.Push(y)
	}
	if auxStack.Size() == 1 {
		stack.Push(auxStack.Peek())
	}
	return result
}

func main() {
	stack := NewStack(1)
	stack.Push(4)
	stack.Push(5)
	stack.Push(-2)
	stack.Push(-3)
	stack.Push(11)
	stack.Push(10)
	stack.Push(5)
	stack.Push(6)
	stack.Push(20)
	fmt.Println(pairWiseConsecutive(stack))
}
