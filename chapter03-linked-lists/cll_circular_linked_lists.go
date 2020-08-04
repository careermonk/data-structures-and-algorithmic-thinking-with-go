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

// CLL ... A linked list with len, head
type CLL struct {
	size int
	head *CLLNode
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
		size: 0,
		head: nil,
	}
}

func (cll *CLL) Length() int {
	current := cll.head
	count := 1
	if current == nil {
		return 0
	}
	current = current.next
	for current != cll.head {
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
		cll.head = newNode
		cll.head.next = newNode
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
		current := cll.head

		// insert on current
		newNode.next = current

		// traverse this list until we find it's end
		for {
			if current.next == cll.head {
				break
			}
			current = current.next
		}
		current.next = newNode
		cll.head = newNode
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
		current := cll.head
		for {
			if current.next == cll.head {
				break
			}
			current = current.next
		}
		current.next = newNode
		newNode.next = cll.head
		cll.size++
	}
}

// Insert ... Insert in the list on a specific location
func (cll *CLL) Insert(data int, pos int) {
	if !(cll.CheckIfEmptyAndAdd(data)) {
		current := cll.head
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
			if current.next == nil && pos-1 > count {
				break
			}
			if count == pos-1 {
				newNode.next = current.next
				current.next = newNode
				cll.size++
				break
			}
			current = current.next
			count++
		}
	}
}

// CheckIfEmpty ... check if empty
func (cll *CLL) CheckIfEmpty() bool {
	if cll.size == 0 {
		return true
	}
	return false
}

// DeleteBeginning ... Delete from the head of the list
func (cll *CLL) DeleteBeginning() int {
	//check if list if empty
	if !(cll.CheckIfEmpty()) {
		current := cll.head
		deletedElem := current.data
		if cll.size == 1 {
			cll.head = nil
			cll.size--
			return deletedElem
		}
		prevStart := cll.head
		cll.head = current.next
		// traverse till end and update last node's next to updated head
		for {
			if current.next == prevStart {
				break
			}
			current = current.next
		}
		current.next = cll.head
		cll.size--
		return deletedElem
	}
	return -1
}

// DeleteEnd ... Delete the last element from circular list
func (cll *CLL) DeleteEnd() int {
	if !(cll.CheckIfEmpty()) {
		current := cll.head
		deletedEle := current.data
		if cll.size == 1 {
			// delete from beginning
			deletedEle = cll.DeleteBeginning()
			return deletedEle
		}
		//traverse till end
		for {
			if current.next.next == cll.head {
				deletedEle = current.next.data
				break
			}
			current = current.next
		}
		// update last element's next pointer
		current.next = cll.head
		cll.size--
		return deletedEle
	}
	return -1
}

// Delete ... delete an element from circular list
func (cll *CLL) Delete(pos int) int {
	if !(cll.CheckIfEmpty()) {
		current := cll.head
		deletedEle := current.data
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
				deletedEle = current.next.data
				break
			}
			current = current.next
		}
		current.next = current.next.next
		cll.size--
		return deletedEle
	}
	return -1
}

// Display ... Print Circular list
func (cll *CLL) Display() {
	current := cll.head
	for i := 0; i < cll.size; i++ {
		fmt.Print(current.data)
		fmt.Print("-->")
		current = current.next
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
