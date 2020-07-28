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
	"math/rand"
)

type (
	TSTNode struct {
		key             byte
		value           interface{}
		left, eq, right *TSTNode
	}

	TernarySearchTree struct {
		length int
		root   *TSTNode
	}
)

// Create a new ternary search tst
func New() *TernarySearchTree {
	tst := &TernarySearchTree{}
	tst.Init()
	return tst
}

// Initialize the tst (reset it so that it's empty). New will do this for you.
func (this *TernarySearchTree) Init() {
	this.length = 0
	this.root = nil
}

// Get the value at the specified key. Returns nil if not found.
func (this *TernarySearchTree) Get(key string) interface{} {
	if this.length == 0 {
		return nil
	}

	node := this.root
	letters := []byte(key)
	for i := 0; i < len(letters); {
		b := letters[i]
		if b > node.key {
			if node.right == nil {
				return nil
			}
			node = node.right
		} else if b < node.key {
			if node.left == nil {
				return nil
			}
			node = node.left
		} else {
			i++
			if i < len(letters) {
				if node.eq == nil {
					return nil
				}
				node = node.eq
			} else {
				break
			}
		}
	}
	return node.value
}

// Test to see whether or not the given key is contained in the tst.
func (this *TernarySearchTree) Has(key string) bool {
	return this.Get(key) != nil
}

// Insert a new key value pair into the collection
func (this *TernarySearchTree) Insert(key string, value interface{}) {
	// If the value is nil then remove this key from the collection
	if value == nil {
		this.Remove(key)
		return
	}

	if this.length == 0 {
		this.root = &TSTNode{0, nil, nil, nil, nil}
	}

	t := this.root
	letters := []byte(key)
	for i := 0; i < len(letters); {
		b := letters[i]
		if b > t.key {
			if t.right == nil {
				t.right = &TSTNode{b, nil, nil, nil, nil}
			}
			t = t.right
		} else if b < t.key {
			if t.left == nil {
				t.left = &TSTNode{b, nil, nil, nil, nil}
			}
			t = t.left
		} else {
			i++
			if i < len(letters) {
				if t.eq == nil {
					t.eq = &TSTNode{letters[i], nil, nil, nil, nil}
				}
				t = t.eq
			}
		}
	}

	if t.value == nil {
		this.length++
	}
	t.value = value
}

// Get the number of items stored in the tst
func (this *TernarySearchTree) Len() int {
	return this.length
}

// Remove a key from the collection
func (this *TernarySearchTree) Remove(key string) interface{} {
	if this.length == 0 {
		return nil
	}

	var remove *TSTNode
	var direction int

	t := this.root
	letters := []byte(key)
	for i := 0; i < len(letters); {
		b := letters[i]
		if b > t.key {
			// Not in the collection
			if t.right == nil {
				return nil
			}
			// This is a branch so we have to keep it
			remove = t
			direction = 1
			// Move to the next node
			t = t.right
		} else if b < t.key {
			// Not in the collection
			if t.left == nil {
				return nil
			}
			// This is a branch so we have to keep it
			remove = t
			direction = -1
			// Move to the next node
			t = t.left
		} else {
			i++
			if i < len(letters) {
				// Not in the collection
				if t.eq == nil {
					return nil
				}
				// Has a value so we need to keep at least this much
				if t.value != nil {
					remove = t
					direction = 0
				}
				// Move to the next node
				t = t.eq
			}
		}
	}

	// If this was the only item in the tst, set the root pointer to nil
	if this.length == 1 {
		this.root = nil
	} else {
		if direction == -1 {
			remove.left = nil
		} else if direction == 0 {
			remove.eq = nil
		} else {
			remove.right = nil
		}
	}
	this.length--
	return t.value
}
func (this *TernarySearchTree) String() string {
	if this.length == 0 {
		return "{}"
	}

	return this.root.String()
}

// Dump the tst to a string for easier debugging
func (this *TSTNode) String() string {
	str := "{" + string(this.key)
	if this.value != nil {
		str += ":" + fmt.Sprint(this.value)
	}
	if this.left != nil {
		str += this.left.String()
	} else {
		str += " "
	}
	if this.eq != nil {
		str += this.eq.String()
	} else {
		str += " "
	}
	if this.right != nil {
		str += this.right.String()
	} else {
		str += " "
	}
	str += "}"
	return str
}

func randomString() string {
	n := 3 + rand.Intn(10)
	letters := make([]byte, n)
	for i := 0; i < n; i++ {
		letters[i] = byte(97 + rand.Intn(25))
	}
	return string(letters)
}
func main() {
	tst := New()
	tst.Insert("test", 1)
	if tst.Len() != 1 {
		panic("expecting len 1")
	}
	if !tst.Has("test") {
		panic("expecting to find key=test")
	}

	tst.Insert("testing", 2)
	tst.Insert("abcd", 0)

	found := false
	found = tst.Has("testing")
	if !found {
		panic("expecting iterator to find test")
	}
	
	if tst.Get("testing") != 2 {
		panic("expecting value=1")
	}

	tst.Remove("testing")
	tst.Remove("abcd")

	v := tst.Remove("test")
	if tst.Len() != 0 {
		panic("expecting len 0")
	}
	if tst.Has("test") {
		panic("expecting not to find key=test")
	}
	if v.(int) != 1 {
		panic("expecting value=1")
	}
}
