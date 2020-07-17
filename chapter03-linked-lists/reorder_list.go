// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

func reorderList(head *ListNode) {
	if head == nil || head.next == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow, fast = slow.next, fast.next.next
	}

	var prev *ListNode
	for slow != nil {
		slow.next, prev, slow = prev, slow, slow.next
	}

	first := head
	for prev.next != nil {
		first.next, first = prev, first.next
		prev.next, prev = first, prev.next
	}
}
