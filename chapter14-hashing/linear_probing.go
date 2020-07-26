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

import "fmt"

type HashTable struct {
	Size   int
	Keys   []interface{}
	Values []interface{}
}

func CreateHashTable() *HashTable {
	h := &HashTable{}
	h.Size = 7
	h.Keys = make([]interface{}, h.Size)
	h.Values = make([]interface{}, h.Size)
	return h
}

func (h *HashTable) hashFunction(key int) int {
	return key % h.Size
}

func (h *HashTable) Get(key int) int {
	i := h.hashFunction(key)
	for ; h.Keys[i] != nil; i = (i + 1) % h.Size {
		if h.Keys[i] == key {
			return h.Values[i].(int)
		}
	}
	return -1
}

func (h *HashTable) Put(key, value int) {
	i := h.hashFunction(key)
	for ; h.Keys[i] != nil; i = (i + 1) % h.Size {
		if h.Keys[i] == key {
			h.Values[i] = value
			break
		}
	}
	h.Keys[i] = key
	h.Values[i] = value
}

// Display: Print the hashtable in a legible format (publicly callable)
func (h *HashTable) Display() {
	fmt.Printf("----------%d elements-------\n", h.Size)
	for i := 0; i < h.Size; i++ {
		fmt.Println(h.Keys[i], ": ",h.Values[i])
	}
}

func main() {
	h := CreateHashTable()
	h.Display()
	h.Put(1, 2)
	h.Display()
	h.Put(2, 3)
	h.Display()
	h.Put(3, 4)
	h.Display()
	h.Put(4, 5)
	h.Display()
	h.Put(5, 6)
	h.Display()
}
