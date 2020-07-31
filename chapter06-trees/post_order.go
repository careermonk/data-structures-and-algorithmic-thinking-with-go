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

// PostOrder traverses a tree in pre-order
func PostOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%d ", root.data)
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

func main() {
	t1 := NewBinaryTree(10, 1)

	// PostOrder walk with print statements
	PostOrder(t1)
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
