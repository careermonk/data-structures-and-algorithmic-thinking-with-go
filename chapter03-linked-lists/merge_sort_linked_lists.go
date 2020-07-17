// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

func sortList(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	slow, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		slow, fast = slow.next, fast.next.next
	}

	firstTail := slow
	slow = slow.next
	firstTail.next = nil // divide the first list and the second

	first, second := sortList(head), sortList(slow)
	return merge(first, second)
}

func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	curHead := &ListNode{}
	tmpHead := curHead

	for head1 != nil && head2 != nil {
		if head1.data < head2.data {
			curHead.next = head1
			head1 = head1.next
			curHead = curHead.next
		} else {
			curHead.next = head2
			head2 = head2.next
			curHead = curHead.next
		}
	}

	if head1 != nil {
		curHead.next = head1
	} else if head2 != nil {
		curHead.next = head2
	}
	return tmpHead.next
}
