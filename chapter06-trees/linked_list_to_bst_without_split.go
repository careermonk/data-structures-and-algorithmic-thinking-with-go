// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
//

var nums *ListNode

func helper(left, right int) *BSTNode {
	if left > right {
		return nil
	}
	mid := int(left + (right-left)/2)
	leftNode := helper(left, mid-1)
	root := new(TreeNode)
	root.data = nums.data
	root.left = leftNode
	nums = nums.Next
	root.right = helper(mid+1, right)
	return root
}

func SortedListToBST(head *ListNode) *BSTNode {
	nums = head
	n := 0
	for head != nil {
		n++
		head = head.Next
	}
	return helper(0, n-1)
}
