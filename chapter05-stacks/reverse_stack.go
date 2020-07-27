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
	"errors"
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

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	temp := stack.array[stack.top]
	return temp, nil
}

// Drain removes all elements that are currently in the stack.
func (stack *Stack) Drain() {
	stack.array = nil
	stack.top = -1
}

func (stack *Stack) reverseStack() {
	if stack.IsEmpty() {
		return
	}

	data := stack.Pop()
	stack.reverseStack()
	stack.insertAtBottom(data)
}

func (stack *Stack) insertAtBottom(data interface{}) {
	if stack.IsEmpty() {
		stack.Push(data)
		return
	}

	temp := stack.Pop()
	stack.insertAtBottom(data)
	stack.Push(temp)
}

func main() {
	stack := NewStack(1)

	stack.Push(10)
	fmt.Println("Pushed to stack : 10")
	stack.Push(20)
	fmt.Println("Pushed to stack : 20")
	stack.Push(30)
	fmt.Println("Pushed to stack : 30")
	fmt.Println("Stack size: ", stack.Size())

	data, _ := stack.Peek()
	fmt.Println("Top element : ", data)

	stack.reverseStack()
	data, _ = stack.Peek()
	fmt.Println("Top element : ", data)

	data = stack.Pop()
	fmt.Println("Popped from stack : ", data)
	data = stack.Pop()
	fmt.Println("Popped from stack : ", data)
	data = stack.Pop()
	fmt.Println("Popped from stack : ", data)

	stack.Drain()
	fmt.Println("Stack size: ", stack.Size())
}
