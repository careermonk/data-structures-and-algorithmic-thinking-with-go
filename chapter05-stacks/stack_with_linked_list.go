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
import "fmt"

type Stack struct {
	top  *ListNode
	size int
}

type ListNode struct {
	data interface{} // All types satisfy the empty interface, so we can store anything here.
	next *ListNode
}

// Return the stack's length
func (s *Stack) length() int {
	return s.size
}

// push a new element onto the stack
func (s *Stack) push(data interface{}) {
	s.top = &ListNode{data, s.top}
	s.size++
}

// Remove the top element from the stack and return it's data
// If the stack is empty, return nil
func (s *Stack) pop() (data interface{}) {
	if s.size > 0 {
		data, s.top = s.top.data, s.top.next
		s.size--
		return
	}
	return nil
}

// Return true if size of the stack is zero otherwise, false
func (s *Stack) isEmpty() bool{
	return s.size == 0
}


// Return false as it is a linked list implementation.
func (s *Stack) isFull() bool{
	return false
}

// Return the top element's data
// If the stack is empty, return nil
func (s *Stack) peek() (data interface{}) {
	if s.size > 0 {
		data = s.top.data
		return
	}
	return nil
}

func main() {
	stack := new(Stack)
	stack.push("Go")
	fmt.Println(stack.peek().(string))
	stack.push("Data")
	stack.push("Algorithms")
	fmt.Println(stack.peek().(string))
	for stack.length() > 0 {
		// We have to do a type assertion because we get back a variable of type
		// interface{} while the underlying type is a string.
		fmt.Printf("%s ", stack.pop().(string))
	}
	fmt.Println()
}
