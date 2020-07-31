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

// A BinraryTreeNode is a binary tree with integer values.
type BinraryTreeNode struct {
	left  *BinraryTreeNode
	data  int
	right *BinraryTreeNode
}

// PreOrder traverses a tree in pre-order
func PreOrder(root *BinraryTreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	PreOrder(root.left)
	PreOrder(root.right)
}

// PreOrder traverses a tree in pre-order, sending each data on a channel.
func PreOrderWalk(root *BinraryTreeNode, ch chan int) {
	if root == nil {
		return
	}
	ch <- root.data
	PreOrderWalk(root.left, ch)
	PreOrderWalk(root.right, ch)
}

// Walker launches Walk in a new goroutine, and returns a read-only channel of values.
func PreOrderWalker(root *BinraryTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		PreOrderWalk(root, ch)
		close(ch)
	}()
	return ch
}

// NewBinraryTree returns a new, random binary tree holding the values 1k, 2k, ..., nk.
func NewBinraryTree(n, k int) *BinraryTreeNode {
	var root *BinraryTreeNode
	for _, v := range rand.Perm(n) {
		root = insert(root, (1+v)*k)
	}
	return root
}

func insert(root *BinraryTreeNode, v int) *BinraryTreeNode {
	if root == nil {
		return &BinraryTreeNode{nil, v, nil}
	}
	if v < root.data {
		root.left = insert(root.left, v)
		return root
	}
	root.right = insert(root.right, v)
	return root
}

func main() {
	t1 := NewBinraryTree(10, 1)

	// Pre-Order walk with print statements
	PreOrder(t1)
	fmt.Println()
	// Pre-Order walk with channel
	c := PreOrderWalker(t1)
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("%d ", v)
	}
}
