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

import (
	"bytes"
	"errors"
	"fmt"
)

type ListNode struct {
	data interface{}
	next *ListNode
}

// Queue implements a queue backed by a linked list
type Queue struct {
	front *ListNode
	rear  *ListNode
	size  int
}

// enQueue adds an element to the end of the queue.
func (q *Queue) enQueue(data interface{}) {
	rear := new(ListNode)
	rear.data = data

	if q.isEmpty() {
		q.front = rear
	} else {
		oldLast := q.rear
		oldLast.next = rear
	}
	q.rear = rear
	q.size++
}

// deQueue removes the front element from the queue.
func (q *Queue) deQueue() (interface{}, error) {
	if q.isEmpty() {
		q.rear = nil
		return nil, errors.New("unable to dequeue element, queue is empty")
	}

	data := q.front.data
	q.front = q.front.next
	q.size--
	return data, nil
}

// front returns the front element in the queue without removing it.
func (q *Queue) frontElement() (interface{}, error) {
	if q.isEmpty() {
		return nil, errors.New("unable to peek element, queue is empty")
	}
	return q.front.data, nil
}

// isEmpty returns whether the queue is empty or not.
func (q *Queue) isEmpty() bool {
	return q.front == nil
}

// size returns the number of elements in the queue.
func (q *Queue) length() int{
	return q.size
}

// String returns a string representation of queue q formatted
// from front to rear.
func (q *Queue) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	j := q.front
	for i := 0; i < q.size; i++ {
		result.WriteString(fmt.Sprintf("%v", j.data))
		if i < q.size-1 {
			result.WriteByte(' ')
		}
		j = j.next
	}
	result.WriteByte(']')
	return result.String()
}

func main() {
	q :=  new(Queue)
	q.enQueue(1)
	q.enQueue(6)
	q.enQueue(7)
	q.enQueue(8)
	q.enQueue(9)
	q.enQueue(2)
	q.enQueue(3)
	q.enQueue(4)
	q.enQueue(5)
	q.enQueue(6)
	q.enQueue(7)
	q.enQueue(8)
	q.enQueue(9)
	fmt.Println(q.String())
	want := "[1 6 7 8 9 2 3 4 5 6 7 8 9]"
	if s := q.String(); s != want {
		fmt.Println("q.String() = ", s, "want = ", want)
	}

	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	want = "[]"
	fmt.Println(q.String())
	if s := q.String(); s != want {
		fmt.Println("q.String() = ", s, "want = ", want)
	}
}
