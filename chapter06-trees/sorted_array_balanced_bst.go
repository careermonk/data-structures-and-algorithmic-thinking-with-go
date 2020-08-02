// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
//

func SortedArrayToBST(A []int) *BSTNode {
	if A == nil {
		return nil
	}
	return helper(A, 0, len(A)-1)
}

func helper(A []int, low int, high int) *BSTNode {
	if low > high { // Done
		return nil
	}
	mid := low + (high-low)/2
	node := new(BSTNode)
	node.data = A[mid]
	node.left = helper(A, low, mid-1)
	node.right = helper(A, mid+1, high)
	return node
}
