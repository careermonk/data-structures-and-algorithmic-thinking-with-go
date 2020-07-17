// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

func findLoopLength(head *ListNode) int {
	fast, slow := head, head
	loopExists := false
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			loopExists = true
			break
		}
	}
	if loopExists {
		counter := 1
		fast = fast.next
		for slow != fast {
			fast = fast.next
			counter++
		}
		return counter
	}
	return 0
}
