// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.


func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    h := reverse(head)
    head.next = nil
    return h
}

func reverse(current *ListNode) *ListNode {
    if current == nil {
        return nil
    }
    temp := reverse(current.next)
    if temp == nil {
       return current
    }
    current.next.next = current
    return temp
}
