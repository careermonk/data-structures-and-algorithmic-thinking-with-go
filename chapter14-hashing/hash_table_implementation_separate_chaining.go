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

var (
	minLoadFactor    = 0.25
	maxLoadFactor    = 0.75
	defaultTableSize = 3
)

type Element struct {
	key   int
	value int
	next  *Element
}

type Hash struct {
	buckets []*Element
}

type HashTable struct {
	table *Hash
	size  *int
}

// createHashTable: Called by checkLoadFactorAndUpdate when creating a new hash, for internal use only.
func createHashTable(tableSize int) HashTable {
	num := 0
	hash := Hash{make([]*Element, tableSize)}
	return HashTable{table: &hash, size: &num}
}

// CreateHashTable: Called by the user to create a hashtable.
func CreateHashTable() HashTable {
	num := 0
	hash := Hash{make([]*Element, defaultTableSize)}
	return HashTable{table: &hash, size: &num}
}

// hashFunction: Used to calculate the index of record within the slice
func hashFunction(key int, size int) int {
	return key % size
}

// Display: Print the hashtable in a legible format (publicly callable)
func (h *HashTable) Display() {
	fmt.Printf("----------%d elements-------\n", *h.size)
	for i, node := range h.table.buckets {
		fmt.Printf("%d :", i)
		for node != nil {
			fmt.Printf("[%d, %d]->", node.key, node.value)
			node = node.next
		}
		fmt.Println("nil")
	}
}

// put: inserts a key into the hash table, for internal use only
func (h *HashTable) put(key int, value int) bool {
	index := hashFunction(key, len(h.table.buckets))
	iterator := h.table.buckets[index]
	node := Element{key, value, nil}
	if iterator == nil {
		h.table.buckets[index] = &node
	} else {
		prev := &Element{0, 0, nil}
		for iterator != nil {
			if iterator.key == key { // Key already exists
				iterator.value = value
				return false
			}
			prev = iterator
			iterator = iterator.next
		}
		prev.next = &node
	}
	*h.size += 1
	return true
}

// Put: inserts a key into the hash table (publicly callable)
func (h *HashTable) Put(key int, value int) {
	sizeChanged := h.put(key, value)
	if sizeChanged == true {
		h.checkLoadFactorAndUpdate()
	}
}

// Get: Retrieve a value for a key from the hash table (publicly callable)
func (h *HashTable) Get(key int) (bool, int) {
	index := hashFunction(key, len(h.table.buckets))
	iterator := h.table.buckets[index]
	for iterator != nil {
		if iterator.key == key { // Key already exists
			return true, iterator.value
		}
		iterator = iterator.next
	}
	return false, 0
}

// del: remove a key-value record from the hash table, for internal use only
func (h *HashTable) del(key int) bool {
	index := hashFunction(key, len(h.table.buckets))
	iterator := h.table.buckets[index]
	if iterator == nil {
		return false
	}
	if iterator.key == key {
		h.table.buckets[index] = iterator.next
		*h.size--
		return true
	} else {
		prev := iterator
		iterator = iterator.next
		for iterator != nil {
			if iterator.key == key {
				prev.next = iterator.next
				*h.size--
				return true
			}
			prev = iterator
			iterator = iterator.next
		}
		return false
	}
}

// Del: remove a key-value record from the hash table (publicly available)
func (h *HashTable) Del(key int) bool {
	sizeChanged := h.del(key)
	if sizeChanged == true {
		h.checkLoadFactorAndUpdate()
	}
	return sizeChanged
}

// getLoadFactor: calculate the loadfactor for the hashtable
// Calculated as: number of buckets stored / length of underlying slice used
func (h *HashTable) getLoadFactor() float64 {
	return float64(*h.size) / float64(len(h.table.buckets))
}

// checkLoadFactorAndUpdate: if 0.25 > loadfactor or 0.75 < loadfactor,
// update the underlying slice to have have loadfactor close to 0.5
func (h *HashTable) checkLoadFactorAndUpdate() {
	if *h.size == 0 {
		return
	} else {
		loadFactor := h.getLoadFactor()
		if loadFactor < minLoadFactor {
			fmt.Println("** Loadfactor below limit, reducing hashtable size **")
			hash := createHashTable(len(h.table.buckets) / 2)
			for _, record := range h.table.buckets {
				for record != nil {
					hash.put(record.key, record.value)
					record = record.next
				}
			}
			h.table = hash.table
		} else if loadFactor > maxLoadFactor {
			fmt.Println("** Loadfactor above limit, increasing hashtable size **")
			hash := createHashTable(*h.size * 2)
			for _, record := range h.table.buckets {
				for record != nil {
					hash.put(record.key, record.value)
					record = record.next
				}
			}
			h.table = hash.table
		}
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
	h.Del(1)
	h.Display()
	h.Del(2)
	h.Display()
	h.Del(3)
	h.Display()
	h.Put(3, 4)
	h.Display()
	h.Put(4, 5)
	h.Display()
	h.Put(5, 6)
	h.Display()
	h.Del(4)
	h.Display()
	h.Del(5)
	h.Display()
	h.Put(11, 12)
	h.Display()
	h.Put(12, 13)
	h.Display()
	h.Put(13, 14)
	h.Display()
	h.Put(14, 15)
	h.Display()
	h.Put(15, 16)
	h.Display()
}
