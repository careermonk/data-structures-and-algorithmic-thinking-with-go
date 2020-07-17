// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

func partition(head *ListNode, X int) *ListNode {
	lesser, greater := &ListNode{}, &ListNode{}
	lesserHead, greaterHead := lesser, greater

	for head != nil {
		if head.data < X {
			lesser.next = head
			lesser = lesser.next
		} else {
			greater.next = head
			greater = greater.next
		}
		head = head.next
	}
	lesser.next, greater.next = nil, nil // cut the connection
	if greaterHead.next != nil {
		lesser.next = greaterHead.next
	}
	return lesserHead.next
}
