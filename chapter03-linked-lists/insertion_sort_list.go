// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.
s
func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    result := &ListNode{Next: head}
    cur := head.Next
    head.Next = nil
    
    for cur != nil {
        pre := result
        target := result.Next
        for target != nil && cur.Val > target.Val {
            target = target.Next
            pre = pre.Next
        }
        temp := cur
        cur = cur.Next
        temp.Next = target
        pre.Next = temp
    }
    
    return result.Next
}
