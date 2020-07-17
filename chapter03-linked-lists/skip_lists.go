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
	"fmt"
	"math/rand"
	"time"
)

const MAX_LEVEL = 32

type SkiplistNode struct {
	data  int
	level []*SkiplistNode
}

type Skiplist struct {
	head            *SkiplistNode
	currentMaxLevel int
	random          *rand.Rand
}

func createSkipList() Skiplist {
	skl := Skiplist{
		head:            new(SkiplistNode),
		currentMaxLevel: 0,
	}

	skl.head.level = make([]*SkiplistNode, MAX_LEVEL)
	source := rand.NewSource(time.Now().UnixNano())
	skl.random = rand.New(source)

	return skl
}

func (this *Skiplist) search(target int) bool {
	p := this.head
	for i := this.currentMaxLevel - 1; i >= 0; i-- {
		for p.level[i] != nil {
			if p.level[i].data == target {
				return true
			}
			if p.level[i].data > target {
				break
			}
			p = p.level[i]
		}
	}

	return false
}

func (this *Skiplist) add(num int) {
	update := make([]*SkiplistNode, MAX_LEVEL)
	p := this.head
	for i := this.currentMaxLevel - 1; i >= 0; i-- {
		for p.level[i] != nil && p.level[i].data < num {
			p = p.level[i]
		}
		update[i] = p
	}

	level := 1
	for this.random.Intn(2) == 1 {
		level++
	}
	if level > MAX_LEVEL {
		level = MAX_LEVEL
	}

	if level > this.currentMaxLevel {
		for i := this.currentMaxLevel; i < level; i++ {
			update[i] = this.head
		}
		this.currentMaxLevel = level
	}

	newNode := new(SkiplistNode)
	newNode.data = num
	newNode.level = make([]*SkiplistNode, level)

	for i := 0; i < level; i++ {
		newNode.level[i] = update[i].level[i]
		update[i].level[i] = newNode
	}
}

func (this *Skiplist) remove(num int) bool {
	if this.head.level[0] == nil { //Skiplist is null
		return false
	}

	update := make([]*SkiplistNode, MAX_LEVEL)
	p := this.head
	for i := this.currentMaxLevel - 1; i >= 0; i-- {
		for p.level[i] != nil && p.level[i].data < num {
			p = p.level[i]
		}
		update[i] = p
	}

	if update[0].level[0] == nil || update[0].level[0].data != num {
		return false
	}

	level := len(update[0].level[0].level)
	for i := 0; i < level; i++ {
		update[i].level[i] = update[i].level[i].level[i]
	}

	for i := this.currentMaxLevel - 1; this.head.level[i] == nil; i-- {
		this.currentMaxLevel--
	}

	return true
}
func main() {
	sl := createSkipList()
	sl.add(10)
	sl.add(3)
	sl.add(30)
	sl.add(11)
	sl.add(7)
	sl.add(89)
	sl.add(2)
	sl.add(19)
	sl.add(23)
	sl.add(3)
	sl.add(5)
	sl.add(4)
	sl.add(29)
	sl.add(33)
	sl.add(20)
	sl.add(9)
	sl.add(6)
	sl.add(14)

	fmt.Println(sl.search(66))
	sl.add(66)
	fmt.Println(sl.search(66))
	fmt.Println(sl.search(19))
	fmt.Println(sl.remove(19))
}
