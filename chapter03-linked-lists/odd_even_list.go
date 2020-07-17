// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.next == nil {
		return head
	}

	oddsHead := head
	evensHead := head.next
	for current, temp := head, head.next; temp != nil; current, temp = temp, temp.next {
		current.next = temp.next
	}

	oddsTail := oddsHead
	for ; oddsTail.next != nil; oddsTail = oddsTail.next {
	}
	oddsTail.next = evensHead

	return oddsHead
}
