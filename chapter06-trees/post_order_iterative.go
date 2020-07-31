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
	"math/rand"
)

// A BinaryTreeNode is a binary tree with integer values.
type BinaryTreeNode struct {
	left  *BinaryTreeNode
	data  int
	right *BinaryTreeNode
}

// NewBinaryTree returns a new, random binary tree holding the values 1k, 2k, ..., nk.
func NewBinaryTree(n, k int) *BinaryTreeNode {
	var root *BinaryTreeNode
	for _, v := range rand.Perm(n) {
		root = insert(root, (1+v)*k)
	}
	return root
}

func insert(root *BinaryTreeNode, v int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{nil, v, nil}
	}
	if v < root.data {
		root.left = insert(root.left, v)
		return root
	}
	root.right = insert(root.right, v)
	return root
}

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
	var i interface{}
	if stack.IsEmpty() {
		return i
	}
	temp := stack.array[stack.top]
	stack.top--
	if stack.Size() < stack.capacity/2 {
		stack.Resize()
	}
	return temp
}

func (stack *Stack) Peek() interface{} {
	var i interface{}
	if stack.IsEmpty() {
		return i
	}
	temp := stack.array[stack.top]
	return temp
}

// Drain removes all elements that are currently in the stack.
func (stack *Stack) Drain() {
	stack.array = nil
	stack.top = -1
}

// PostOrderWalk traverses a tree in in-order, sending each data on a channel.
func PostOrderWalk(root *BinaryTreeNode, ch chan int) {
	if root == nil {
		return
	}

	PostOrderWalk(root.left, ch)
	PostOrderWalk(root.right, ch)
	ch <- root.data
}

// PostOrderWalker launches Walk in a new goroutine, and returns a read-only channel of values.
func PostOrderWalker(root *BinaryTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		PostOrderWalk(root, ch)
		close(ch)
	}()
	return ch
}

func PostOrder1(root *BinaryTreeNode) []int {
	var result []int
	stack := NewStack(1)
	stack.Push(root)
	for !stack.IsEmpty() {
		temp := stack.Pop().(*BinaryTreeNode)

		result = append(result, temp.data)

		if temp.left != nil {
			stack.Push(temp.left)
		}

		if temp.right != nil {
			stack.Push(temp.right)
		}
	}
	n := len(result)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		result[i], result[j] = result[j], result[i]
	}
	fmt.Println(result)
	return result
}

func PostOrder2(root *BinaryTreeNode) []int {
	result := []int{}
	stack := []*BinaryTreeNode{root}

	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root == nil {
			continue
		}

		result = append(result, root.data)

		stack = append(stack, root.left)
		stack = append(stack, root.right)

	}
	// reverse
	n := len(result)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		result[i], result[j] = result[j], result[i]
	}
	fmt.Println(result)
	return result
}

func main() {
	t1 := NewBinaryTree(10, 1)
	PostOrder1(t1)
	fmt.Println()
	PostOrder2(t1)
	fmt.Println()

	// PostOrderWalker walk with channel
	c := PostOrderWalker(t1)
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("%d ", v)
	}
}
