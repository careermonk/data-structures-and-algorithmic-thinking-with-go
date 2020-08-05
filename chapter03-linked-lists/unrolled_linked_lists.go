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

var blockSize int //max number of nodes in a block
type ListNode struct {
	data int
	next *ListNode
}

type LinkedBlock struct {
	nodeCount int
	next      *LinkedBlock
	head      *ListNode
}

var blockHead *LinkedBlock

// create an empty block
func newLinkedBlock() *LinkedBlock {
	block := &LinkedBlock{}
	block.next = nil
	block.head = nil
	block.nodeCount = 0
	return block
}

//create a node
func newListNode(data int) *ListNode {
	newNode := &ListNode{}
	newNode.next = nil
	newNode.data = data
	return newNode
}

func searchElement(k int, fLinkedBlock **LinkedBlock, fListNode **ListNode) {
	//find the block
	j := (k + blockSize - 1) / blockSize //k-th node is in the j-th block
	p := blockHead
	j = j - 1
	for j > 0 {
		p = p.next
		j = j - 1
	}
	*fLinkedBlock = p

	//find the node
	q := p.head
	k = k % blockSize
	if k == 0 {
		k = blockSize
	}
	k = p.nodeCount + 1 - k
	for k > 0 {
		k = k - 1
		q = q.next
	}
	*fListNode = q
}

//start shift operation from block *p
func shift(A *LinkedBlock) {
	var B *LinkedBlock
	var temp *ListNode
	for A.nodeCount > blockSize { //if this block still have to shift
		if A.next == nil { //reach the end. A little different
			A.next = newLinkedBlock()
			B = A.next
			temp = A.head.next
			A.head.next = A.head.next.next
			B.head = temp
			temp.next = temp
			A.nodeCount--
			B.nodeCount++
		} else {
			B = A.next
			temp = A.head.next
			A.head.next = A.head.next.next
			temp.next = B.head.next
			B.head.next = temp
			B.head = temp
			A.nodeCount--
			B.nodeCount++
		}
		A = B
	}
}

func addElement(k, x int) {
	var p, q *ListNode
	var r *LinkedBlock

	if blockHead == nil { //initial, first node and block
		blockHead = newLinkedBlock()
		blockHead.head = newListNode(x)
		blockHead.head.next = blockHead.head
		blockHead.nodeCount++
	} else {
		if k == 0 { //special case for k=0.
			p = blockHead.head
			q = p.next
			p.next = newListNode(x)
			p.next.next = q
			blockHead.head = p.next
			blockHead.nodeCount++
			shift(blockHead)
		} else {
			searchElement(k, &r, &p)
			q = p
			for q.next != p {
				q = q.next
			}
			q.next = newListNode(x)
			q.next.next = p
			r.nodeCount++
			shift(r)
		}
	}
}

func search(k int) int {
	var p *ListNode
	var q *LinkedBlock
	searchElement(k, &q, &p)
	return p.data
}

// Test Code
func main() {
	blockSize = 3
	addElement(1, 15)
	addElement(2, 25)
	addElement(3, 35)
	addElement(0, 45)
	fmt.Println(search(1))
	fmt.Println(search(2))
	fmt.Println(search(3))
}
