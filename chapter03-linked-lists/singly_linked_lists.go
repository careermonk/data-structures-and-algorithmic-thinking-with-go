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
	node := &ListNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = node
	} else {
		node.next = ll.head
		ll.head = node
	}
	ll.size++
	return
}

func (ll *LinkedList) insertAtEnd(data interface{}) {
	node := &ListNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	ll.size++
	return
}

// Insert adds an item at position i
func (ll *LinkedList) insert(i int, data interface{}) error {
	if i < 0 || i > ll.size {
		return fmt.Errorf("Index out of bounds")
	}
	node := &ListNode{data, nil}
	if i == 0 {
		node.next = ll.head
		ll.head = node
		return nil
	}
	current := ll.head
	j := 0
	for j < i-2 {
		j++
		current = current.next
	}
	node.next = current.next
	current.next = node
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
func (ll *LinkedList) delete(i int) (interface{}, error) {
	if i < 0 || i > ll.size {
		return nil, fmt.Errorf("Index out of bounds")
	}
	current := ll.head
	j := 0
	for j < i-1 {
		j++
		current = current.next
	}
	remove := current.next
	current.next = remove.next
	ll.size--
	return remove.data, nil
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

func main() {
	ll := LinkedList{}
	fmt.Printf("insertAtBeginning: A\n")
	ll.insertAtBeginning("A")
	fmt.Printf("insertAtBeginning: B\n")
	ll.insertAtBeginning("B")
	fmt.Printf("insertAtEnd: C\n")
	ll.insertAtEnd("C")

	fmt.Printf("length: %d\n", ll.length())

	err := ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("deleteFront\n")
	_, err = ll.deleteFront()
	if err != nil {
		fmt.Printf("deleteFront Error: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("deleteLast\n")
	_, err = ll.deleteLast()
	if err != nil {
		fmt.Printf("deleteLast Error: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("deleteLast\n")
	_, err = ll.deleteLast()
	if err != nil {
		fmt.Printf("deleteLast Error: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("deleteLast\n")
	_, err = ll.deleteLast()
	if err != nil {
		fmt.Printf("deleteLast Error: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("length: %d\n", ll.length())
}
