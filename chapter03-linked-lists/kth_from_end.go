// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

func kthFromEnd(head *ListNode, n int) *ListNode {
    first, second := head, head
    for ; n > 0; n-- {
        second = second.next
    }
    for ; second.next != nil; first, second = first.next, second.next {
    }
    first.next = first.next.next
    return first
}  
