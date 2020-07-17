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
    newHead, _ := reverseList(head.next, n - m + 1)
    head.next = newHead
    // if m == 1 { return head.next } else { return dummy.next }
    return dummy.next
}

// return new head and the head of the rest
func reverseList(head *ListNode, count int) (*ListNode, *ListNode) {
    if count == 1 { return head, head.next }
    newHead, restHead := reverseList(head.next, count - 1)
    head.next.next = head
    head.next = restHead
    return newHead, restHead
}
