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
    // Duplicate each node.
    X := head
    for X != nil {
        newNode := &ListNode{
            data: X.data,
            next: X.next,
            random: X.random,
        }
        
        X.next = newNode
        X = newNode.next
    }
    
    // Make the duplicated nodes point to the correct random.
    X = head
    for X != nil {
        X = X.next
        if X.random != nil {
            X.random = X.random.next
        }
        X = X.next
    }
    
    // Extract the duplicated nodes and recover the original list.
    result := &ListNode{}
    Y := result
    X = head
    for X != nil {
        n := X.next
        X.next = X.next.next
        X = X.next
        Y.next = n
        Y = Y.next
    }
    
    return result.next
}
