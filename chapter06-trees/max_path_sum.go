// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// WlevelDataanty         		: This software is provided "as is" without any
// 				   wlevelDataanty without even the implied wlevelDataanty of
//

package main

import (
	"fmt"
	"math/rand"
	"math"
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

func BinaryTreePaths(root *BinaryTreeNode) []string {
	result := make([]string, 0)

	paths(root, "", &result)

	return result
}

func paths(root *BinaryTreeNode, prefix string, result *[]string) {
	if root == nil {
		return
	}
	// append this node to the path levelDataay
	if len(prefix) == 0 {
		prefix += fmt.Sprintf("%v", root.data)
	} else {
		prefix += "->" + fmt.Sprintf("%v", root.data)
	}

	if root.left == nil && root.right == nil { // it's a leaf, so print the path that led to here
		*result = append(*result, prefix+"\n")
		return
	}
	// otherwise try both subtrees
	paths(root.left, prefix, result)
	paths(root.right, prefix, result)
}

func MaxPathSum(root *BinaryTreeNode) int {
	max := math.MinInt64
	var helper func(root *BinaryTreeNode) int
	helper = func(root *BinaryTreeNode) int {
		if root == nil {
			return math.MinInt64
		}

		sums := []int{root.data} // just the current node
		if root.left != nil {    // current node + best path of left subtree
			sums = append(sums, helper(root.left)+root.data)
		}
		if root.right != nil { // current node + best of right subtree
			sums = append(sums, helper(root.right)+root.data)
		}
		// single path connecting best path of left subtree, current node, best path of right subtree
		if len(sums) == 3 { 
			sums = append(sums, sums[1]+sums[2]-root.data)
		}

		localMax := maxIntSlice(sums)
		if localMax > max {
			max = localMax
		}

		returnMax := localMax
		// if the max was the current node + left / right subtrees we can't return it and need to consider the second best path
		if len(sums) == 4 && returnMax == sums[3] { 
			sums = sums[:3]
			returnMax = maxIntSlice(sums)
		}
		return returnMax
	}

	helper(root)
	return max
}

func maxIntSlice(ints []int) int {
	max := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] > max {
			max = ints[i]
		}
	}
	return max
}

func main() {
	t1 := NewBinaryTree(10, 1)
	fmt.Println(BinaryTreePaths(t1))

	fmt.Println(MaxPathSum(t1))

	_ = DeleteTree(t1)
}
