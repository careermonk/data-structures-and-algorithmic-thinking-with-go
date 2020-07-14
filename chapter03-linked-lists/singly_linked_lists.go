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
func (ll *LinkedList) insert(i int, data interface{}) error {
	newNode := &ListNode{data, nil}
	if i <= 0 {
		newNode.next = ll.head
		ll.head = newNode
		ll.size++
		return nil
	}
	prev := ll.head
	// find the n-1th newNode (predecessor): decrement n and move down the list
	for prev != nil {
		i = i - 1
		if i == 0 {
			break
		}
		prev = prev.next // repoint temp to the next element
	}
	if prev == nil {
		return fmt.Errorf("insert: Index out of bounds")
	}
	newNode.next = prev.next
	prev.next = newNode
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
	k := 1
	var q *ListNode
	if ll.head == nil {
		return nil, fmt.Errorf("deleteLast: List is empty")
	}
	p := ll.head

	if i <= 1 { // from the beginning
		ll.head = ll.head.next
		return p.data, nil
	}

	for (p != nil) && k < i { // traverse the list to the position from which we want to delete
		k++
		q = p
		p = p.next
	}
	if p == nil { /* at the end */
		return nil, fmt.Errorf("Index out of bounds")
	}

	q.next = p.next

	return p.data, nil
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
	fmt.Printf("insert: D\n")
	ll.insert(3, "D")
	fmt.Printf("length: %d\n", ll.length())

	err := ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("delete\n")
	_, err = ll.delete(5)
	if err != nil {
		fmt.Printf("deleteError: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
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
