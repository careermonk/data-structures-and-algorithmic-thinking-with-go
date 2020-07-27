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
	"math"
)

type MinStack struct {
	elementStack  []int
	minimumsStack []int
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// Initialize your data structure here.
func NewMinStack() MinStack {
	return MinStack{}
}

func (stack *MinStack) Push(data int) {
	stack.elementStack = append(stack.elementStack, data)

	if len(stack.minimumsStack) == 0 {
		stack.minimumsStack = append(stack.minimumsStack, data)
	} else {
		minimum := min(stack.minimumsStack[len(stack.minimumsStack)-1], data)
		stack.minimumsStack = append(stack.minimumsStack, minimum)
	}
}

func (stack *MinStack) Pop() int {
	if len(stack.elementStack) > 0 {
		popped := stack.elementStack[len(stack.elementStack)-1]
		stack.elementStack = stack.elementStack[:len(stack.elementStack)-1]
		stack.minimumsStack = stack.minimumsStack[:len(stack.minimumsStack)-1]
		return popped
	} else {
		return math.MaxInt32
	}
}

func (stack *MinStack) Peek() int {
	if len(stack.elementStack) > 0 {
		return stack.elementStack[len(stack.elementStack)-1]
	} else {
		return 0
	}
}

func (stack *MinStack) Size() int {
	return len(stack.elementStack)
}

func (stack *MinStack) GetMin() int {
	if len(stack.minimumsStack) > 0 {
		return stack.minimumsStack[len(stack.minimumsStack)-1]
	} else {
		return 0
	}
}

func (stack *MinStack) IsEmpty() bool {
	return len(stack.elementStack) == 0
}

func (stack *MinStack) Clear() {
	stack.elementStack = nil
	stack.minimumsStack = nil
}

func main() {
	minimumsStack := NewMinStack()

	//Push int values
	minimumsStack.Push(-1)
	minimumsStack.Push(2)
	minimumsStack.Push(0)
	minimumsStack.Push(-4)

	// Stack Size() test
	if minimumsStack.Size() != 4 {
		fmt.Println("Test failed: Expected stack size = 4 from [-1, 2, 0, -4]")
	} else { //Just for printing
		fmt.Println("Test passed: Expected stack size = 4 from [-1, 2, 0, -4]")
	}

	// Stack GetMin() test
	if minimumsStack.GetMin() != -4 {
		fmt.Println("Test failed: Expected get minimum = -4 from [-1, 2, 0, -4]")
	} else { //Just for printing
		fmt.Println("Test passed: Expected get minimum = -4 from [-1, 2, 0, -4]")
	}

	// Stack Top() test
	if minimumsStack.Peek() != -4 {
		fmt.Println("Test failed: Expected top of stack = -4 from [-1, 2, 0, -4]")
	} else { //Just for printing
		fmt.Println("Test passed: Expected top of stack = -4 from [-1, 2, 0, -4]")
	}

	// Pop() test
	minimumsStack.Pop() // [-1, 2, 0]
	if minimumsStack.Peek() != 0 {
		fmt.Println("Test failed: After pop, Expected top of stack = 0 from [-1, 2, 0]")
	} else { //Just for printing
		fmt.Println("Test passed: After pop, Expected top of stack = 0 from [-1, 2, 0]")
	}

	// Stack GetMin() test
	if minimumsStack.GetMin() != -1 {
		fmt.Println("Test failed: Expected get minimum = -1 from [-1, 2, 0]")
	} else { //Just for printing
		fmt.Println("Test passed: Expected get minimum = -1 from [-1, 2, 0]")
	}

	// Clear() and isEmpty() test
	minimumsStack.Clear()
	if minimumsStack.IsEmpty() != true {
		fmt.Println("Test failed: After Clear(), Stack is expected to be empty.")
	} else { //Just for printing
		fmt.Println("Test passed: After Clear(), Stack is expected to be empty.")
	}
}
