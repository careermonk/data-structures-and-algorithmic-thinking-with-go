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

//CLLNode ... newNode of circular linked list
type CLLNode struct {
	data int
	next *CLLNode
}

// CLL ... A linked list with len, start
type CLL struct {
	size  int
	start *CLLNode
}

// NewNode ... Creating a new node
func NewNode(data int) *CLLNode {
	return &CLLNode{
		data: data,
		next: nil,
	}
}

// NewCircularLinkedList ... Creating an empty circular linked list
func NewCircularLinkedList() CLL {
	return CLL{
		size:  0,
		start: nil,
	}
}

// CheckIfEmpty ... check if empty
func (cll *CLL) CheckIfEmpty() bool {
	if cll.size == 0 {
		return true
	}
	return false
}


func (cll *CLL) Length() int {
	current := cll.start
	count := 1
	if current == nil {
		return 0
	}
	current = current.next
	for current != cll.start {
		current = current.next
		count++
	}
	return count
}

// CheckIfEmptyAndAdd ... check if doubly link list is empty
func (cll *CLL) CheckIfEmptyAndAdd(data int) bool {
	newNode := &CLLNode{
		data: data,
		next: nil,
	}
	// check if list is empty
	if cll.size == 0 {
		// insert first node in doubly linked list
		cll.start = newNode
		cll.start.next = newNode
		cll.size++
		return true
	}
	return false
}


// InsertBeginning ... insert in the beginning
func (cll *CLL) InsertBeginning(data int) {
	if !(cll.CheckIfEmptyAndAdd(data)) {
		newNode := &CLLNode{
			data: data,
			next: nil,
		}
		head := cll.start

		// insert on head
		newNode.next = head

		// traverse this list until we find it's end
		for {
			if head.next == cll.start {
				break
			}
			head = head.next
		}
		head.next = newNode
		cll.start = newNode
		cll.size++
	}
}

// InsertEnd ... insert in the end
func (cll *CLL) InsertEnd(data int) {
	if !(cll.CheckIfEmptyAndAdd(data)) {
		newNode := &CLLNode{
			data: data,
			next: nil,
		}
		head := cll.start
		for {
			if head.next == cll.start {
				break
			}
			head = head.next
		}
		head.next = newNode
		newNode.next = cll.start
		cll.size++
	}
}

// InsertLocation ... Insert in the list on a specific location
func (cll *CLL) InsertLocation(data int, pos int) {
	if !(cll.CheckIfEmptyAndAdd(data)) {
		head := cll.start
		count := 1
		if pos == 1 {
			cll.InsertBeginning(data)
			return
		}
		newNode := &CLLNode{
			data: data,
			next: nil,
		}
		for {
			if head.next == nil && pos-1 > count {
				break
			}
			if count == pos-1 {
				newNode.next = head.next
				head.next = newNode
				cll.size++
				break
			}
			head = head.next
			count++
		}
	}
}

// DeleteBeginning ... Delete from the start of the list
func (cll *CLL) DeleteBeginning() int {
	//check if list if empty
	if !(cll.CheckIfEmpty()) {
		head := cll.start
		deletedElem := head.data
		if cll.size == 1 {
			cll.start = nil
			cll.size--
			return deletedElem
		}
		prevStart := cll.start
		cll.start = head.next
		// traverse till end and update last node's next to updated start
		for {
			if head.next == prevStart {
				break
			}
			head = head.next
		}
		head.next = cll.start
		cll.size--
		return deletedElem
	}
	return -1
}

// DeleteEnd ... Delete the last element from circular list
func (cll *CLL) DeleteEnd() int {
	if !(cll.CheckIfEmpty()) {
		head := cll.start
		deletedEle := head.data
		if cll.size == 1 {
			// delete from beginning
			deletedEle = cll.DeleteBeginning()
			return deletedEle
		}
		//traverse till end
		for {
			if head.next.next == cll.start {
				deletedEle = head.next.data
				break
			}
			head = head.next
		}
		// update last element's next pointer
		head.next = cll.start
		cll.size--
		return deletedEle
	}
	return -1
}

// DeleteFromPosition ... delete an element from circular list
func (cll *CLL) DeleteFromPosition(pos int) int {
	if !(cll.CheckIfEmpty()) {
		head := cll.start
		deletedEle := head.data
		if cll.size == 1 {
			// delete from beginning
			deletedEle = cll.DeleteBeginning()
			return deletedEle
		}
		if cll.size == pos {
			// delete from end
			deletedEle = cll.DeleteEnd()
			return deletedEle
		}
		// delete from middle
		//traverse till you find position
		count := 1
		for {
			if count == pos-1 {
				deletedEle = head.next.data
				break
			}
			head = head.next
		}
		head.next = head.next.next
		cll.size--
		return deletedEle
	}
	return -1
}

// Display ... Print Circular list
func (cll *CLL) Display() {
	head := cll.start
	for i := 0; i < cll.size; i++ {
		fmt.Print(head.data)
		fmt.Print("-->")
		head = head.next
	}
	fmt.Println()
}

func main() {
	cll := CLL{}
	fmt.Println("Size: ", cll.Length())
	cll.InsertBeginning(100)
	cll.InsertBeginning(90)
	cll.InsertBeginning(70)
	cll.InsertBeginning(60)
	cll.InsertBeginning(50)
	cll.InsertBeginning(40)
	cll.InsertBeginning(30)

	cll.InsertBeginning(19) // Update dll

	cll.Display()
	fmt.Println("Size: ", cll.Length())
}
