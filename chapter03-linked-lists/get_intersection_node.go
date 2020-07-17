// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

func getIntersectionNode(list1, list2 *LinkedList) *ListNode {
    len1, len2 := findLen(list1), findLen(list2)
    head1, head2 := ll1.head, ll2.head
    if len1 > len2 {
        for ; len1 > len2; len1-- {
            headA = headA.next
        }
    } else {
        for ; len2 > len1; len2-- {
            headB = headB.next
        }
    }
    for headA != headB {
        headA, headB = headA.next, headB.next
    }
    return headA
}
func findLen(ll *LinkedList) int {
    l := 0
    current = ll.head
    for ; current != nil; current = current.next {
        l++
    }
    return l
}
