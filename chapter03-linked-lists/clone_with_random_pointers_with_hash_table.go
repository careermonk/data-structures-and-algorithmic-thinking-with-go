// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

type ListNode struct { // defines a ListNode in a singly linked list
	data   interface{} // the datum
	next   *ListNode   // pointer to the next ListNode
	random *ListNode   // pointer to a random ListNode
}

func clone(head *ListNode) *ListNode {
	var (
		result *ListNode               = &ListNode{}
		Y                              = result
		X      *ListNode               = head
		HT     map[*ListNode]*ListNode = make(map[*ListNode]*ListNode, 0)
	)
	for X != nil {
		Y.next = &ListNode{data: X.data}

		HT[X] = Y.next

		Y = Y.next
		X = X.next
	}

	X = head
	Y = result.next
	for X != nil {
		if n, found := HT[X.random]; found {
			Y.random = n
		}
		Y = Y.next
		X = X.next
	}
	return result.next
}
