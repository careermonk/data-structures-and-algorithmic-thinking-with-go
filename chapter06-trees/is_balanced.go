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

const NotBalanced = -1

const NotBalanced = -1

func isBalanced(root *BinarySearchTreeNode) bool {
	if root == nil {
		return true
	}
	return checkHeight(root) != NotBalanced
}

func checkHeight(root *BinarySearchTreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := checkHeight(root.Left)
	if leftHeight == NotBalanced {
		return NotBalanced
	}

	rightHeight := checkHeight(root.Right)
	if rightHeight == NotBalanced {
		return NotBalanced
	}

	if abs(leftHeight-rightHeight) > 1 {
		return false
	}

	return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
