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
	data  interface{}
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
		var level []interface{}
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

func LevelOrder(root *BinaryTreeNode) [][]interface{} { // Data from each level is being returned as a separate list
	var result [][]interface{}
	if root == nil {
		return result
	}
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		var level []interface{}
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

// PreOrder traverses a tree in pre-order
func PreOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	PreOrder(root.left)
	PreOrder(root.right)
}

// InOrder traverses a tree in pre-order
func InOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	InOrder(root.left)
	fmt.Printf("%d ", root.data)
	InOrder(root.right)
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

func BinaryTreePaths(root *BinaryTreeNode) []string {
	result := make([]string, 0)

	paths(root, "", &result)

	return result
}

func paths(root *BinaryTreeNode, prefix string, result *[]string) {
	if root == nil {
		return
	}
	// append this node to the path array
	if len(prefix) == 0 {
		prefix += fmt.Sprintf("%v", string(root.data.(uint8)))
	} else {
		prefix += "->" + fmt.Sprintf("%v", string(root.data.(uint8)))
	}

	if root.left == nil && root.right == nil { // it's a leaf, so print the path that led to here
		*result = append(*result, prefix+"\n")
		return
	}
	// otherwise try both subtrees
	paths(root.left, prefix, result)
	paths(root.right, prefix, result)
}

func BuildTreeFromPreOrder(preOrder string, i *int) *BinaryTreeNode {
	newNode := &BinaryTreeNode{
		data: preOrder[*i],
	}

	if len(preOrder) == 0 { // Boundary Condition
		return nil
	}
	if preOrder[*i] == 'L' { // On reaching leaf node, return
		return newNode
	}

	*i = *i + 1 // populate left sub tree
	newNode.left = BuildTreeFromPreOrder(preOrder, i)

	*i = *i + 1 // populate right sub tree
	newNode.right = BuildTreeFromPreOrder(preOrder, i)

	return newNode
}

func main() {
	i := 0
	t1 := BuildTreeFromPreOrder("ILILL", &i)
	fmt.Println(BinaryTreePaths(t1))

	t1 = DeleteTree(t1)
}
