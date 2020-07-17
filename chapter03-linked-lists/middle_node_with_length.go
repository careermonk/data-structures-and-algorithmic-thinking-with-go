// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

func middleNode(head *ListNode) *ListNode {
    l := length(head)
    count, target := 0, (l/2)+1
    
    for {
        count++
        if count == target {
            return head
        }
        head = head.Next
    }
}

func length(head *ListNode) int {
    current, count := head, 0

    for current != nil {
        count++
        current = current.Next
    }
    
    return count
}
