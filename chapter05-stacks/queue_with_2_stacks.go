// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
//

package main

import (
	"fmt"
)

type Queue struct {
	stack1 []int
	stack2 []int
}

func NewQueue() Queue {
	return Queue{}
}

// Push element data to the back of queue.
func (q *Queue) EnQueue(data int) {
	q.stack1 = append(q.stack1, data)
}

// Removes the element from in front of queue and returns that element.
func (q *Queue) DeQueue() int {
	if len(q.stack2) == 0 {
		for len(q.stack1) > 0 {
			item := q.stack1[len(q.stack1)-1]
			q.stack1 = q.stack1[:len(q.stack1)-1]
			q.stack2 = append(q.stack2, item)
		}
	}

	item := q.stack2[len(q.stack2)-1]
	q.stack2 = q.stack2[:len(q.stack2)-1]

	return item
}

// Get the front element.
func (q *Queue) Front() int {
	if len(q.stack2) < 1 {
		for len(q.stack1) > 0 {
			item := q.stack1[len(q.stack1)-1]
			q.stack1 = q.stack1[:len(q.stack1)-1]
			q.stack2 = append(q.stack2, item)
		}
	}

	return q.stack2[len(q.stack2)-1]
}

// Returns whether the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.stack1) == 0 && len(q.stack2) == 0
}

func main() {
	obj := NewQueue()
	obj.EnQueue(10)
	obj.EnQueue(20)
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.Front())	// prints 20
	fmt.Println(obj.IsEmpty())
	obj.EnQueue(30)
	obj.EnQueue(40)
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.Front())	// prints 30
	fmt.Println(obj.IsEmpty())
}
