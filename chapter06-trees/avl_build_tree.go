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
	"strconv"
)

type AVLTreeNode struct {
	data   int
	left   *AVLTreeNode
	right  *AVLTreeNode
	height int
}

func DeleteTree(root *AVLTreeNode) *AVLTreeNode {
	// free left child
	if root.left != nil {
		root.left = DeleteTree(root.left)
	}
	// free right child
	if root.right != nil {
		root.right = DeleteTree(root.right)
	}
	// a leaf, just return nil, a garbage collector will free the memory
	return nil
}

func printSpace(spaceNum int) {
	for i := 0; i < spaceNum; i++ {
		fmt.Printf("  ")
	}
}

// PreOrder traverses a root in pre-order
func PreOrder(root *AVLTreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	PreOrder(root.left)
	PreOrder(root.right)
}

// InOrder traverses a root in pre-order
func InOrder(root *AVLTreeNode) {
	if root == nil {
		return
	}

	InOrder(root.left)
	fmt.Printf("%d ", root.data)
	InOrder(root.right)
}

// PostOrder traverses a root in pre-order
func PostOrder(root *AVLTreeNode) {
	if root == nil {
		return
	}

	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%d ", root.data)
}

func LevelOrder(root *AVLTreeNode) [][]int { // Data from each level is being returned as a separate list
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	queue := []*AVLTreeNode{root}
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

func BinaryTreePaths(root *AVLTreeNode) []string {
	result := make([]string, 0)

	paths(root, "", &result)

	return result
}

func paths(root *AVLTreeNode, prefix string, result *[]string) {
	if root == nil {
		return
	}
	// append this node to the path array
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

// find element position in the avl tree
func (root *AVLTreeNode) Find(x int) *AVLTreeNode {
	if root != nil {
		if x < root.data {
			// find in left
			return root.left.Find(x)
		} else if x > root.data {
			// find right
			return root.right.Find(x)
		} else if x == root.data {
			// we find that
			return root
		}
	}

	//can not found or empty root
	return nil
}

// find min element in the tree
func FindMin(root *AVLTreeNode) *AVLTreeNode {
	if root == nil {
		return nil
	}
	for root.left != nil {
		root = root.left
	}
	return root
}

// find max element in the tree
func FindMax(root *AVLTreeNode) *AVLTreeNode {
	if root == nil {
		return nil
	}
	for root.right != nil {
		root = root.right
	}
	return root
}

// return the height of the node
func Height(node *AVLTreeNode) int {
	if node == nil {
		return -1
	} else {
		return node.height
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Insert an element x to AvlTree
func Insert(root *AVLTreeNode, x int) *AVLTreeNode {
	if root == nil {
		//root = new(AVLTreeNode)
		root = &AVLTreeNode{
			data:   x,
			height: 0,
		}
	} else if x < root.data {
		root.left = Insert(root.left, x)
		if Height(root.left)-Height(root.right) == 2 {
			if x < root.left.data { //left left
				root = singleLeftRotate(root)
			} else { //left right
				root = doubleRotateLeftRight(root)
			}
		}
	} else if x > root.data {
		root.right = Insert(root.right, x)
		if Height(root.left)-Height(root.right) == 2 {
			if x > root.right.data {
				root = singleRightRotate(root)
			} else {
				root = doubleRotateRightLeft(root)
			}
		}
	} else {
		//x existed, we do nothing.
	}

	root.height = max(Height(root.left), Height(root.right)) + 1

	return root
}

// Delete an element from AvlTree
func Delete(root *AVLTreeNode, x int) *AVLTreeNode {
	var tmp *AVLTreeNode
	if root != nil {
		if x < root.data {
			root.left = Delete(root.left, x)
			if Height(root.right)-Height(root.left) == 2 {
				if root.right.right != nil {
					root = singleRightRotate(root)
				} else {
					root = doubleRotateRightLeft(root)
				}
			} else {
				root.height = max(Height(root.left), Height(root.right)) + 1
			}
		} else if x > root.data {
			root.right = Delete(root.right, x)
			if Height(root.left)-Height(root.right) == 2 {
				if root.left.left != nil {
					root = singleLeftRotate(root)
				} else {
					root = doubleRotateLeftRight(root)
				}

			} else {
				root.height = max(Height(root.left), Height(root.right)) + 1
			}
		} else if root.left != nil && root.right != nil {
			tmp = FindMin(root.right)
			root.data = tmp.data
			root.right = Delete(root.right, tmp.data)

		} else {
			if root.left == nil {
				root = root.right
			} else if root.right == nil {
				root = root.left
			}
		}
	}

	return root
}

// left rotate a root, and update node's height, return the new root
func singleLeftRotate(X *AVLTreeNode) *AVLTreeNode {
	var W *AVLTreeNode
	if X != nil {
		W = X.left
		X.left = W.right
		W.right = X

		// update height
		X.height = max(Height(X.left), Height(X.right)) + 1
		W.height = max(Height(W.left), Height(W.right)) + 1

		X = W
	}
	return X
}

// right rotate a root, and update node's height, return the new root
func singleRightRotate(W *AVLTreeNode) *AVLTreeNode {
	var X *AVLTreeNode
	if W != nil {
		X = W.right
		W.right = X.left
		X.left = W

		// update height
		W.height = max(Height(W.left), Height(W.right)) + 1
		X.height = max(Height(X.left), Height(X.right)) + 1

		W = X
	}
	return W
}

func doubleRotateLeftRight(Z *AVLTreeNode) *AVLTreeNode {
	Z.left = singleRightRotate(Z.left)

	return singleLeftRotate(Z)
}

func doubleRotateRightLeft(Z *AVLTreeNode) *AVLTreeNode {
	Z.right = singleLeftRotate(Z.right)

	return singleRightRotate(Z)
}

var count = 1

func BuildHB0(h int) *AVLTreeNode {
	if h == 0 {
		return nil
	}
	newNode := &AVLTreeNode{}
	newNode.left = BuildHB0(h - 1)
	newNode.data = count //assume count is a global variable
	count++
	newNode.right = BuildHB0(h - 1)
	return newNode
}

func main() {
	t1 := BuildHB0(4)
	fmt.Println(LevelOrder(t1))
	fmt.Println(BinaryTreePaths(t1))

	fmt.Print("InOrder Traversal: ")
	InOrder(t1)
	fmt.Printf("\nmin: %v, max:%v\n", FindMin(t1).data, FindMax(t1).data)
}
