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

type Stack struct {
	Q1 Queue
	Q2 Queue
}

// Initialize your data structure here.
func NewStack() Stack {
	return Stack{}
}

// Push element data onto stack.
func (stack *Stack) Push(data int) {
	if stack.Q1.IsEmpty() { // Insert to queue whichever is not empty
		stack.Q2.EnQueue(data)
	} else {
		stack.Q1.EnQueue(data)
	}
}

// Removes the element on top of the stack and returns that element.
func (stack *Stack) Pop() int {
	if stack.Q2.IsEmpty() {
		i, n := 1, stack.Q1.Size()
		for i < n && !stack.Q1.IsEmpty() { // Transfer n-1 elements
			stack.Q2.EnQueue(stack.Q1.DeQueue())
			i++
		}
		return stack.Q1.DeQueue() // Delete the front element
	}

	i, n := 1, stack.Q2.Size()
	for i < n && !stack.Q2.IsEmpty() { // Transfer n-1 elements
		stack.Q1.EnQueue(stack.Q2.DeQueue()) // Delete the front element
		i++
	}
	return stack.Q2.DeQueue()
}

// Get the top element.
func (stack *Stack) Peek() int {
	if stack.Q2.IsEmpty() {
		i, n := 1, stack.Q1.Size()
		for i < n && !stack.Q1.IsEmpty() {
			stack.Q2.EnQueue(stack.Q1.DeQueue())
			i++
		}
		temp := stack.Q1.DeQueue() // The last element is the top element in the stack
		stack.Q2.EnQueue(temp)
		return temp
	}
	i, n := 1, stack.Q2.Size()
	for i < n && !stack.Q2.IsEmpty() {
		stack.Q1.EnQueue(stack.Q2.DeQueue())
		i++
	}
	temp := stack.Q2.DeQueue() // The last element is the top element in the stack
	stack.Q1.EnQueue(temp)
	return temp
}

// Returns whether the stack is empty.
func (stack *Stack) IsEmpty() bool {
	return stack.Q1.IsEmpty() && stack.Q2.IsEmpty()
}

type Queue struct {
	Data []int
}

func (queue *Queue) EnQueue(data int) {
	queue.Data = append(queue.Data, data)
}

func (queue *Queue) Front() int {
	if len(queue.Data) > 0 {
		return queue.Data[0]
	}
	return 0
}

func (queue *Queue) DeQueue() int {
	if len(queue.Data) > 0 {
		data := queue.Data[0]
		queue.Data = queue.Data[1:]
		return data
	}
	return 0
}

func (queue *Queue) Size() int {
	return len(queue.Data)
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.Data) == 0
}

func main() {
	obj := NewStack()
	obj.Push(10)
	obj.Push(20)
	obj.Push(30)
	fmt.Println(obj.Peek()) // prints 30
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek()) // prints 20
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek()) // prints 20
	fmt.Println(obj.Pop())

}
