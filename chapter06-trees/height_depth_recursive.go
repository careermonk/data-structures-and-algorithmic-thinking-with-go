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

func Insert(root *BinaryTreeNode, v int) *BinaryTreeNode {
	newNode := &BinaryTreeNode{nil, v, nil}
	if root == nil {
		return newNode
	}
	if root.left == nil {
		root.left = Insert(root.left, v)
	} else if root.right == nil {
		root.right = Insert(root.right, v)
	} else {
		randomize := rand.Intn(1) // select either left or right randomly; It generate 0 or 1
		if randomize == 0 {
			root.left = Insert(root.left, v)
		} else {
			root.right = Insert(root.right, v)
		}
	}
	return root
}

// Compute the number of nodes in a tree.
func Size(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	} else {
		return Size(root.left) + 1 + Size(root.right)
	}
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

// Compute the height (or depth) of a tree
func Height(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	} else {
		// compute the depth of each subtree
		leftheight := Height(root.left)
		rightheight := Height(root.right)

		if leftheight > rightheight {
			return leftheight + 1
		} else {
			return rightheight + 1
		}
	}
}

func main() {
	t1 := NewBinaryTree(20, 1)
	fmt.Println(LevelOrder(t1))
	fmt.Println(Size(t1))
	fmt.Println(Height(t1))

	t1 = DeleteTree(t1)

	fmt.Println(LevelOrder(t1))
	fmt.Println(Size(t1))
	fmt.Println(Height(t1))
}
