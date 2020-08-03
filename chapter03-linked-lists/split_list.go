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

type ListNode struct { // defines a ListNode in a linked list
	data interface{} // the datum
	next *ListNode   // pointer to the next ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func (ll *LinkedList) insertAtBeginning(data interface{}) {
	newNode := &ListNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
	ll.size++
	return
}

func (ll *LinkedList) insertAtEnd(data interface{}) {
	newNode := &ListNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.size++
	return
}

// Insert adds an item at position i
func (ll *LinkedList) insert(position int, data interface{}) error {
	// This condition to check whether the position given is valid or not.
	if position < 1 || position > ll.size+1 {
		return fmt.Errorf("insert: Index out of bounds")
	}
	newNode := &ListNode{data, nil}

	var prev, current *ListNode
	prev = nil
	current = ll.head

	for position > 1 {
		prev = current
		current = current.next
		position = position - 1
	}

	if prev != nil {
		prev.next = newNode
		newNode.next = current
	} else {
		newNode.next = current
		ll.head = newNode
	}

	ll.size++
	return nil
}

func (ll *LinkedList) deleteFront() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}
	data := ll.head.data
	ll.head = ll.head.next
	ll.size--
	return data, nil
}

func (ll *LinkedList) deleteLast() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteLast: List is empty")
	}
	var prev *ListNode
	current := ll.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		ll.head = nil
	}
	ll.size--
	return current.data, nil
}

// delete removes an element at position i
func (ll *LinkedList) delete(position int) (interface{}, error) {
	// This condition to check whether the position given is valid or not.
	if position < 1 || position > ll.size+1 {
		return nil, fmt.Errorf("insert: Index out of bounds")
	}
	// Complete this method
	var prev, current *ListNode
	prev = nil
	current = ll.head

	pos := 0

	if position == 1 {
		ll.head = ll.head.next
	} else {
		for pos != position-1 {
			pos = pos + 1
			prev = current
			current = current.next
		}

		if current != nil {
			prev.next = current.next
		}
	}
	ll.size--
	return current.data, nil
}

func (ll *LinkedList) display() error {
	if ll.head == nil {
		return fmt.Errorf("display: List is empty")
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println()
	return nil
}

// with extra size field
func (ll *LinkedList) length() int {
	return ll.size
}

func display(head *ListNode) error {
	if head == nil {
		return fmt.Errorf("display: List is empty")
	}
	current := head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println()
	return nil
}

// length returns the linked list size
func (ll *LinkedList) length2() int {
	size := 0
	current := ll.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

// Uses the fast/slow pointer strategy
func splitList(head *ListNode) (head1 *ListNode, head2 *ListNode) {
	var slow, fast *ListNode
	if head == nil || head.next == nil { // length < 2 cases
		head1 = head
		head2 = nil
	} else {
		slow = head
		fast = head.next
		// Advance 'fast' two nodes, and advance 'slow' one ListNode
		for fast != nil {
			fast = fast.next
			if fast != nil {
				slow = slow.next
				fast = fast.next
			}
		}
		// 'slow' is before the midpoint in the list, so split it in two at that point.
		head1 = head
		head2 = slow.next
		slow.next = nil
	}
	return head1, head2
}

func main() {
	list1 := LinkedList{}
	list1.insertAtBeginning(100)
	list1.insertAtBeginning(90)
	list1.insertAtBeginning(70)
	list1.insertAtBeginning(60)
	list1.insertAtBeginning(50)
	list1.insertAtBeginning(40)
	list1.insertAtBeginning(30)
	list1.insertAtBeginning(96)
	list1.insertAtBeginning(64)
	list1.insertAtBeginning(58)
	list1.insertAtBeginning(29)
	err := list1.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	l1, l2 := splitList(list1.head)

	err = display(l1)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = display(l2)
	if err != nil {
		fmt.Println(err.Error())
	}
}
