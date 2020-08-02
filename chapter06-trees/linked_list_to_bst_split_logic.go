 // Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
//

func SplitList(head *ListNode) (*ListNode, *ListNode, *ListNode) {

	if head.next == nil {
		return nil, head, nil
	}

	slowPointer := head
	fastPointer := head
	previousPointer := head
	for fastPointer != nil && fastPointer.next != nil {
		previousPointer = slowPointer
		fastPointer = fastPointer.next.next
		slowPointer = slowPointer.next
	}
	previousPointer.next = nil

	return head, slowPointer, slowPointer.next
}

func SortedListToBST(head *ListNode) *BSTNode {
	if head == nil {
		return nil
	}

	left, middle, right := SplitList(head)

	var node BSTNode
	node.data = middle.data
	node.left = SortedListToBST(left)
	node.Right = SortedListToBST(right)
	return &node
}
