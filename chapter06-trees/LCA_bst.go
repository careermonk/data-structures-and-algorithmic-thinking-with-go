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
	"strconv"
)

// A BSTNode is a binary search tree with integer values.
type BSTNode struct {
	left  *BSTNode
	data  int
	right *BSTNode
}

// NewBinarySearchTree returns a new, random binary search tree holding the values 1k, 2k, ..., nk.
func NewBinarySearchTree(n, k int) *BSTNode {
	var root *BSTNode
	for _, v := range rand.Perm(n) {
		root = Insert(root, (1+v)*k)
	}
	return root
}

func Insert(root *BSTNode, v int) *BSTNode {
	if root == nil {
		return &BSTNode{nil, v, nil}
	}
	if v < root.data {
		root.left = Insert(root.left, v)
		return root
	}
	root.right = Insert(root.right, v)
	return root
}

func Delete(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return nil
	}

	if data < root.data {
		root.left = Delete(root.left, data)
	} else if data > root.data {
		root.right = Delete(root.right, data)
	} else {

		if root.right == nil {
			return root.left
		}

		if root.left == nil {
			return root.right
		}

		t := root
		root = findMin(t.right)
		root.right = deleteMin(t.right)
		root.left = t.left
	}
	return root
}

func deleteMin(root *BSTNode) *BSTNode {
	if root.left == nil {
		return root.right
	}

	root.left = deleteMin(root.left)

	return root
}

func DeleteTree(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	}
	// first delete both subtrees
	root.left = DeleteTree(root.left)
	root.right = DeleteTree(root.right)
	// Delete current root only after deleting subtrees
	root = nil
	return root
}

// PreOrder traverses a tree in pre-order
func PreOrder(root *BSTNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	PreOrder(root.left)
	PreOrder(root.right)
}

// InOrder traverses a tree in pre-order
func InOrder(root *BSTNode) {
	if root == nil {
		return
	}

	InOrder(root.left)
	fmt.Printf("%d ", root.data)
	InOrder(root.right)
}

// PostOrder traverses a tree in pre-order
func PostOrder(root *BSTNode) {
	if root == nil {
		return
	}

	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%d ", root.data)
}

func LevelOrder(root *BSTNode) [][]int { // Data from each level is being returned as a separate list
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	queue := []*BSTNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			root := queue[0]
			level = append(level, root.data)
			queue = queue[1:]

			if root.left != nil {
				queue = append(queue, root.left)
			}
			if root.right != nil {
				queue = append(queue, root.right)
			}
		}
		result = append(result, level)
	}
	return result
}

func BinaryTreePaths(root *BSTNode) []string {
	result := make([]string, 0)

	paths(root, "", &result)

	return result
}

func paths(root *BSTNode, prefix string, result *[]string) {
	if root == nil {
		return
	}
	// append this root to the path array
	if len(prefix) == 0 {
		prefix += strconv.Itoa(root.data)
	} else {
		prefix += "->" + strconv.Itoa(root.data)
	}

	if root.left == nil && root.right == nil { // it's a leaf, so print the path that led to here
		*result = append(*result, prefix+"\n")
		return
	}
	// otherwise try both subtrees
	paths(root.left, prefix, result)
	paths(root.right, prefix, result)
}

func findMin(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	}
	for root.left != nil {
		root = root.left
	}
	return root

}

func findMax(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	}
	for root.right != nil {
		root = root.right
	}
	return root
}

func LCA(root *BSTNode, α, β int) *BSTNode {
	cur := root
	for {
		switch {
		case α < cur.data && β < cur.data:
			cur = cur.left
		case α > cur.data && β > cur.data:
			cur = cur.right
		default:
			return cur
		}
	}
	return root
}

func main() {
	t1 := NewBinarySearchTree(15, 1)
	fmt.Println(LevelOrder(t1))
	fmt.Println(BinaryTreePaths(t1))

	fmt.Println(LCA(t1, 4, 15))
}
