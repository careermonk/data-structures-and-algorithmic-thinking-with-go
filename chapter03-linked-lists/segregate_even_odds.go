// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

package main

type ListNode struct { // defines a ListNode in a singly linked list
	data int       // the datum
	next *ListNode // pointer to the next ListNode
}

// Function to separate even and odd nodes
func segregateEvenOdds(head *ListNode) *ListNode {
	var evensHead, evenEnd, oddsHead, oddEnd *ListNode
	evensHead, evenEnd, oddsHead, oddsHead = nil, nil, nil, nil

	// ListNode to traverse the list.
	currNode := head

	for currNode != nil {
		val := currNode.data
		// If current data is even, add it to evens list
		if val%2 == 0 {
			if evensHead == nil {
				evensHead = currNode
				evenEnd = evensHead
			} else {
				evenEnd.next = currNode
				evenEnd = evenEnd.next
			}
		} else { // If current data is odd, add it to odds list
			if oddsHead == nil {
				oddsHead = currNode
				oddEnd = oddsHead
			} else {
				oddEnd.next = currNode
				oddEnd = oddEnd.next
			}
		}

		// Move head pointer one step in forward direction
		currNode = currNode.next
	}

	// If either odd list or even list is empty, no change is required as all elements are either even or odd
	if oddsHead == nil || evensHead == nil {
		return head
	}

	// Add odd list after even list
	evenEnd.next = oddsHead
	oddEnd.next = nil

	return evensHead
}
