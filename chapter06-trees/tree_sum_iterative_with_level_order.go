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
		root = Insert(root, (1+v)*k)
	}
	return root
}

func Insert(root *BinaryTreeNode, v int) *BinaryTreeNode {
	newNode := &BinaryTreeNode{nil, v, nil}
	if root == nil {
		return newNode
	}
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			} else {
				node.left = newNode
				return root
			}
			if node.right != nil {
				queue = append(queue, node.right)
			} else {
				node.right = newNode
				return root
			}
		}
	}
	return root
}

// Compute the number of nodes in a tree.
func Size(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	var result int
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			result++
			level = append(level, node.data)
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}

	}
	return result
}

func LevelOrder(root *BinaryTreeNode) [][]int { // Data from each level is being returned as a separate list
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			level = append(level, node.data)
			queue = queue[1:]

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		result = append(result, level)
	}
	return result
}

func DeleteTree(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	// first delete both subtrees
	root.left = DeleteTree(root.left)
	root.right = DeleteTree(root.right)
	// Delete current node only after deleting subtrees
	root = nil
	return root
}

func Sum(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	var result int
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			result += node.data
			level = append(level, node.data)
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return result
}

func main() {
	t1 := NewBinaryTree(5, 1)
	fmt.Println(LevelOrder(t1))

	fmt.Println(Sum(t1))

	t1 = DeleteTree(t1)
}
