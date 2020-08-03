// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

func getIntersectionNode(head1, head2 *ListNode) *ListNode {
    len1, len2 := findLen(list1), findLen(list2)
    if len1 > len2 {
        for ; len1 > len2; len1-- {
            head1 = head1.next
        }
    } else {
        for ; len2 > len1; len2-- {
            head2 = head2.next
        }
    }
    for head1 != head2 {
        head1, head2 = head1.next, head2.next
    }
    return head1
}
func findLen(head *ListNode) int {
    l := 0
    for ; head != nil; head = head.next {
        l++
    }
    return l
}
