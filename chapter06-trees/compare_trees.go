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

// Walk traverses a tree depth-first, sending each data on a channel.
func Walk(root *BinaryTreeNode, ch chan int) {
	if root == nil {
		return
	}
	Walk(root.left, ch)
	ch <- root.data
	Walk(root.right, ch)
}

// Walker launches Walk in a new goroutine, and returns a read-only channel of values.
func Walker(root *BinaryTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(root, ch)
		close(ch)
	}()
	return ch
}

// Compare reads values from two Walkers that run simultaneously, and returns true if t1 and t2 have the same contents.
func Compare(t1, t2 *BinaryTreeNode) bool {
	c1, c2 := Walker(t1), Walker(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

// NewBinraryTree returns a new, random binary tree holding the values 1k, 2k, ..., nk.
func NewBinraryTree(n, k int) *BinaryTreeNode {
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
	t1 := NewBinraryTree(100, 1)
	fmt.Println(Compare(t1, NewBinraryTree(100, 1)), "Same Contents")
	fmt.Println(Compare(t1, NewBinraryTree(99, 1)), "Differing Sizes")
	fmt.Println(Compare(t1, NewBinraryTree(100, 2)), "Differing Values")
	fmt.Println(Compare(t1, NewBinraryTree(101, 2)), "Dissimilar")
}
