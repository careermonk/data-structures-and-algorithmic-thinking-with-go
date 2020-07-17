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
