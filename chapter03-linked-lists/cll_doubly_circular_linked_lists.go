// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

package main

import "fmt"

type CLLNode struct {
	data int
	next *CLLNode
	prev *CLLNode
}
type CLL struct {
	head *CLLNode
	tail *CLLNode
}

// to avoid mistakes when using pointer vs struct for new CLLNode creation
func newNode(data int) *CLLNode {
	n := &CLLNode{}
	n.data = data
	n.next = nil
	n.prev = nil
	return n
}
func (cll *CLL) insertAtBegin(data int) {
	n := newNode(data)
	n.next = cll.head
	cll.head = n
	if cll.tail == nil {
		cll.tail = n
	}
	n.prev = cll.tail
	cll.tail.next = n
}
func (cll *CLL) insertAtEnd(data int) {
	n := newNode(data)
	if cll.head == nil {
		cll.head = n
		cll.tail = n
		return
	}
	if cll.head == cll.tail {
		cll.head.next = n
		cll.head.prev = n
		n.next = cll.head
		n.prev = cll.head
		cll.tail = n
		return
	}
	cur := cll.head
	for ; cur.next != cll.tail; cur = cur.next {
	}
	cur.next.next = n
	n.prev = cur.next
	n.next = cll.head
	cll.tail = n
}
func (cll *CLL) deleteFrontNode() int {
	if cll.head == nil {
		return -1
	}
	if cll.head == cll.tail {
		value := cll.head
		cll.head = nil
		cll.tail = nil
		return value.data
	}
	cur := cll.head
	cll.head = cur.next
	if cll.head != nil {
		cll.head.prev = cll.tail
		cll.tail.next = cll.head
	}
	return cur.data
}
func (cll *CLL) deleteLastNode() int {
	// no item
	if cll.head == nil {
		return -1
	}
	// only one item
	if cll.head == cll.tail {
		return cll.deleteFrontNode()
	}
	// more than one, go to second last
	cur := cll.head
	for ; cur.next != cll.tail; cur = cur.next {
	}
	retval := cur.next.data
	cur.next = cll.head
	cll.tail = cur
	cll.head.prev = cur
	return retval
}
func (cll *CLL) count() int {
	var ctr int = 0
	for cur := cll.head; cur != cll.tail; cur = cur.next {
		ctr++
	}
	return ctr
}
func (cll *CLL) reverse() {
	var prev, next *CLLNode
	if cll.head == cll.tail {
		return
	}
	cur := cll.head.next
	cll.tail = cll.head
	if cll.head.next != nil {
		for cur != cll.head {
			next = cur.next
			cur.next = prev
			cur.prev = next
			prev = cur
			cur = next
		}
	}
	cll.head = prev
	cur.next = cll.tail.next
	cll.tail.next.next = cll.tail
	cll.tail.next = cll.head
	cll.head.prev = cll.tail
}

func (cll *CLL) display() {
	if cll.head != nil {
		fmt.Print(cll.head.data, " ")
		if cll.head.next != nil {
			for cur := cll.head.next; cur != cll.head; cur = cur.next {
				fmt.Print(cur.data, " ")
			}
		}
	}
	fmt.Print(" ")
}
func (cll *CLL) displayReverse() {
	if cll.head == nil {
		return
	}
	var cur *CLLNode
	for cur = cll.head; cur.next != cll.tail; cur = cur.next {
	}
	for ; cur != cll.tail; cur = cur.prev {
		fmt.Print(cur.data, " ")
	}
	fmt.Print(" ")
}

func main() {
	cll := CLL{}
	cll.insertAtBegin(100)
	cll.insertAtBegin(90)
	cll.insertAtBegin(70)
	cll.insertAtBegin(60)
	cll.insertAtBegin(50)
	cll.insertAtBegin(40)
	cll.insertAtBegin(30)

	cll.insertAtBegin(19) // Update cll

	cll.display()
}
