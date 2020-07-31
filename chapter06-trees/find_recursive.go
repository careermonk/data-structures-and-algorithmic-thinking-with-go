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

func find(root *BinraryTreeNode, data int) *BinraryTreeNode {
	// Base case == empty tree, in that case, the data is not found so return false
	if root == nil {
		return root
	} else {
		//see if found here
		if data == root.data {
			return root
		} else {
			// otherwise recur down the correct subtree
			temp := find(root.left, data)
			if temp != nil {
				return temp
			} else {
				return find(root.right, data)
			}
		}
	}
}

func main() {
	t1 := NewBinraryTree(10, 1)
	fmt.Println(find(t1, 6))
	fmt.Println(find(t1, 19))
}
