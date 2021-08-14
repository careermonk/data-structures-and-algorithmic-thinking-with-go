// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

package main

import "fmt"

// DLLNode ... newNode of doubly linked list
type DLLNode struct {
	data int
	prev *DLLNode // points to previous node in link list
	next *DLLNode // points to next node in link list
}

// DLL ...
type DLL struct {
	size int
	head *DLLNode
	tail *DLLNode
}

// NewNode ... it return a new node
func NewNode(data int) *DLLNode {
	return &DLLNode{
		data: data,
		prev: nil,
		next: nil,
	}
}

// NewDLL ... returns a doubly linked list
func NewDLL() DLL {
	return DLL{
		size: 0,
		head: nil,
		tail: nil,
	}
}

//CheckIfEmpty ... check if empty
func (dll *DLL) CheckIfEmpty() bool {
	if dll.size == 0 {
		return true
	}
	return false
}

// CheckIfEmptyAndAdd ... check if doubly link list is empty
func (dll *DLL) CheckIfEmptyAndAdd(newNode *DLLNode) bool {
	// check if list is empty
	if dll.size == 0 {
		// insert first node in doubly linked list
		dll.head = newNode
		dll.tail = newNode
		dll.size++
		return true
	}
	return false
}

// InsertBeginning ... insert in the beginning of doubly linked list
func (dll *DLL) InsertBeginning(data int) {
	newNode := &DLLNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if !(dll.CheckIfEmptyAndAdd(newNode)) {
		head := dll.head
		// update newnode links - prev and next
		newNode.next = head
		newNode.prev = nil

		// update head node
		head.prev = newNode

		// update dll start and length
		dll.head = newNode
		dll.size++
		return
	}
	return
}

// InsertEnd ... inserts a node in the end of doubly linked list
func (dll *DLL) InsertEnd(data int) {
	newNode := &DLLNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if !(dll.CheckIfEmptyAndAdd(newNode)) {
		tail := dll.tail

		// update newnode links - prev and next
		newNode.prev = tail
		newNode.next = nil

		// update head node
		tail.next = newNode

		// update dll end and length
		dll.tail = newNode
		dll.size++
	}
	return
}

// InsertMiddle ... insert between two nodes in doubly linkedlist
func (dll *DLL) InsertMiddle(data int, loc int) {
	newNode := &DLLNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if !(dll.CheckIfEmptyAndAdd(newNode)) {
		head := dll.head
		for i := 1; i < dll.size; i++ {
			if i == loc {
				// update newnode links - prev and next
				newNode.prev = head.prev
				newNode.next = head

				// update head node
				head.prev.next = newNode
				head.prev = newNode

				//keep traversing till we find the location
				dll.size++
				return
			}
			head = head.next
		}
	}
	return
}

// DeleteFirst ... Delete last element
func (dll *DLL) DeleteFirst() int {
	if !(dll.CheckIfEmpty()) {
		head := dll.head
		deletedNode := head.data

		if dll.head != dll.tail {
			// update doubly linked list
			dll.head = head.next
			dll.head.prev = nil
		} else {
			dll.head = nil
			dll.tail = nil
		}
		dll.size--

		return deletedNode
	}
	return -1
}

// DeleteLast ... deletes last element from doubly linked list
func (dll *DLL) DeleteLast() int {
	if !(dll.CheckIfEmpty()) {
		tail := dll.tail
		deletedNode := tail.data

		if dll.head != dll.tail {
			// update doubly linked list
			dll.tail = tail.prev
			dll.tail.next = nil
		} else {
			dll.head = nil
			dll.tail = nil
		}
		dll.size--

		return deletedNode
	}
	return -1
}

// DeleteMiddle .. delete from any location
func (dll *DLL) DeleteMiddle(pos int) int {
	if !(dll.CheckIfEmpty()) {
		//delete from any position
		head := dll.head
		for i := 1; i <= pos; i++ {
			if head.next == nil && pos > i {
				//list is lesser than given position
				return -1
			} else if i == pos {
				// delete from this location
				head.prev.next = head.next
				head.next.prev = head.prev
				dll.size--
				return head.data
			}
			head = head.next
		}
	}
	return -1
}

// Display ... Display the doubly linked list
func (dll *DLL) Display() {
	head := dll.head
	for i := 0; i < dll.size; i++ {
		fmt.Println("-----newNode no------", i+1)
		fmt.Println("data", head.data)
		fmt.Println("prev", head.prev)
		fmt.Println("next", head.next)
		fmt.Println("------------------")
		head = head.next
	}
}

func main() {
	dll := DLL{}
	dll.InsertBeginning(100)
	dll.InsertBeginning(90)
	dll.InsertBeginning(70)
	dll.InsertBeginning(60)
	dll.InsertBeginning(50)
	dll.InsertBeginning(40)
	dll.InsertBeginning(30)

	dll.InsertBeginning(19) // Update dll

	dll.Display()

	dll.InsertEnd(119) // Update dll
	dll.Display()

	dll.DeleteFirst() // Update dll
	dll.DeleteLast()  // Update dll
	dll.Display()
}
