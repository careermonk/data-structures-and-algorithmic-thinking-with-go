// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry, result := 0, new(ListNode)
    for node := result; l1 != nil || l2 != nil || carry > 0; node = node.next {
        if l1 != nil {
            carry += l1.data
            l1 = l1.next
        }
        if l2 != nil {
            carry += l2.data
            l2 = l2.next
        }
        node.next = &ListNode{carry%10, nil}
        carry /= 10
    }
    return result.next
}
