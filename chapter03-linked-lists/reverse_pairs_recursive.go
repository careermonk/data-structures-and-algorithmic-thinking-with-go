// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

func reversePairs(head *ListNode, m int, n int) *ListNode {
    dummy := new(ListNode)
    head, dummy.next = dummy, head
    for i := 0; i < m-1; i++ { 
        head = head.next
         }
    var curr, prev *ListNode = head.next, nil
    for i := 0; i < n - m + 1; i++ {
        next := curr.next
        curr.next = prev
        prev = curr
        curr = next
    }
    head.next.next = curr
    head.next = prev
    return dummy.next
}
