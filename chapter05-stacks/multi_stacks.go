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
	"errors"
	"fmt"
	"math"
)

type MultiStacks struct {
	top1, top2 int
	capacity   int
	array      []int
}

// Returns an initialized stacks
func (stacks *MultiStacks) Init(capacity int) *MultiStacks {
	stacks.top1 = -1
	stacks.top2 = capacity
	stacks.capacity = capacity
	stacks.array = make([]int, capacity)
	return stacks
}

// Returns an new stacks
func NewStack(capacity int) *MultiStacks {
	return new(MultiStacks).Init(capacity)
}

// Returns the size of the size
func (stacks *MultiStacks) Size(stackNumber int) int {
	if stackNumber == 1 {
		return stacks.top1 + 1
	} else {
		return stacks.capacity - stacks.top2
	}
}

// MultiStacks is full when top is equal to the last index
func (stacks *MultiStacks) IsFull() bool {
	return (stacks.Size(1) + stacks.Size(1)) == stacks.capacity
}

// MultiStacks is empty when top is equal to -1
func (stacks *MultiStacks) IsEmpty(stackNumber int) bool {
	if stackNumber == 1 {
		return stacks.top1 == -1
	} else {
		return stacks.top2 == stacks.capacity
	}

}

func (stacks *MultiStacks) Push(stackNumber int, data int) error {
	if stacks.IsFull() {
		return errors.New("stacks is full")
	}
	if stackNumber == 1 {
		stacks.top1++
		stacks.array[stacks.top1] = data
	} else {
		stacks.top2 = stacks.top2 - 1
		stacks.array[stacks.top2] = data
	}
	return nil
}

func (stacks *MultiStacks) Pop(stackNumber int) int {
	var result int
	if stacks.IsEmpty(stackNumber) {
		return math.MinInt32
	}
	if stackNumber == 1 {
		result = stacks.array[stacks.top1]
		stacks.top1--
	} else {
		result = stacks.array[stacks.top2]
		stacks.top2++
	}
	return result
}

func (stacks *MultiStacks) Peek(stackNumber int) int {
	var result int
	if stacks.IsEmpty(stackNumber) {
		return math.MinInt32
	}
	if stackNumber == 1 {
		result = stacks.array[stacks.top1]
	} else {
		result = stacks.array[stacks.top2]
	}
	return result
}

// Drain removes all elements that are currently in the stacks.
func (stacks *MultiStacks) Drain() {
	stacks.array = nil
	stacks.top1 = -1
	stacks.top2 = stacks.capacity
}

func main() {
	stacks := NewStack(100)

	stacks.Push(1, 10)
	fmt.Println("Pushed to stacks : 10")
	stacks.Push(1, 20)
	fmt.Println("Pushed to stacks : 20")
	stacks.Push(2, 30)
	fmt.Println("Pushed to stacks : 30")
	stacks.Push(2, 40)
	fmt.Println("Pushed to stacks : 40")
	fmt.Println("MultiStacks size: ", stacks.Size(1), stacks.Size(2))

	fmt.Println("Top elements : ", stacks.Peek(1), stacks.Peek(2))

	fmt.Println("Popped from stacks : ", stacks.Pop(1), stacks.Pop(2))
	fmt.Println("MultiStacks size: ", stacks.Size(1), stacks.Size(2))
	fmt.Println("Top elements : ", stacks.Peek(1), stacks.Peek(2))

	stacks.Drain()
	fmt.Println("MultiStacks size: ", stacks.Size(1), stacks.Size(2))
}
