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

// Trie Holder
type Trie struct {
	letter   rune
	children []*Trie
	meta     map[string]interface{}
	isLeaf   bool
}

// NewTrie creates a new Trie and initializes it with default options
func NewTrie() *Trie {
	newTrie := &Trie{}
	newTrie.children = []*Trie{}
	newTrie.meta = make(map[string]interface{})

	return newTrie
}

func (trie *Trie) hasChild(a rune) (bool, *Trie) {
	for _, child := range trie.children {
		if child.letter == a {
			return true, child
		}
	}

	return false, nil
}

func (trie *Trie) addChild(a rune) *Trie {
	nw := NewTrie()
	nw.letter = a
	//nw.parent = trie
	trie.children = append(trie.children, nw)
	return nw
}

// Count returns the number of words in a trie.
func (trie *Trie) Count() int {
	count := 0
	for _, child := range trie.children {
		if child.isLeaf == true {
			count++
		}

		count += child.Count()
	}

	return count
}

// Add a word to a trie
func (trie *Trie) Add(word string) *Trie {
	letters, node, i := []rune(word), trie, 0
	n := len(letters)

	for i < n {
		if exists, value := node.hasChild(letters[i]); exists {
			node = value
		} else {
			node = node.addChild(letters[i])
		}

		i++

		if i == n {
			node.isLeaf = true
		}
	}

	return node
}

// Remove a word from a trie.
// This does not free memory, it just marks the node.
func (trie *Trie) Remove(word string) {
	a := trie.Find(word)

	if a != nil {
		a.isLeaf = false
	}
}

// FindNode returns the node whether it is a word or not.
func (trie *Trie) FindNode(word string) *Trie {
	letters, node, i := []rune(word), trie, 0
	n := len(letters)

	for i < n {
		if exists, value := node.hasChild(letters[i]); exists {
			node = value
		} else {
			return nil
		}

		i++
	}

	return node
}

// Find returns the node pointing to a word
func (trie *Trie) Find(word string) *Trie {
	node := trie.FindNode(word)

	if node == nil {
		return nil
	}

	if node.isLeaf != true {
		return nil
	}

	return node
}

// Get metadata belonging to a node.
func (trie *Trie) Get(key string) (interface{}, bool) {
	if trie == nil {
		return nil, false
	}

	if _, ok := trie.meta[key]; ok {
		return trie.meta[key], true
	}

	return nil, false
}

// Set metadata for a word
func (trie *Trie) Set(key string, value interface{}) {
	if trie == nil {
		return
	}
	trie.meta[key] = value
}
func main() {
	rootTrie := NewTrie()
	rootTrie.Add("hello")
	rootTrie.Add("heart")

	if rootTrie.Find("hello") == nil {
		panic("a problem")
	}

	if rootTrie.Find("heart") == nil {
		panic("a problem")
	}
	if rootTrie.Find("hearts") != nil {
		panic("a problem")
	}

	rootTrie.Add("hearts")

	if rootTrie.Find("hearts") == nil {
		panic("a problem")
	}

	if rootTrie.FindNode("he") == nil {
		panic("a problem")
	}

	rootTrie.Add("tester")

	if rootTrie.FindNode("testing") != nil {
		panic("a problem")
	}

	if rootTrie.FindNode("te") == nil {
		panic("a problem")
	}

	if rootTrie.FindNode("test") == nil {
		panic("a problem")
	}

	rootTrie = NewTrie()
	rootTrie.Add("ab")
	rootTrie.Add("abc")
	rootTrie.Add("abcd")
	rootTrie.Add("abcdefg")

	rootTrie.Add("zoidberg")
	rootTrie.Add("iceberg")

	if rootTrie.Count() != 6 {
		panic("a problem")
	}

	rootTrie.Add("heart")

	node := rootTrie.Find("heart")

	node.Set("test", "alpha")
	item, exists := node.Get("test")

	if exists != true {
		panic("a problem")
	}

	if item.(string) != "alpha" {
		panic("a problem")
	}
	rootTrie.Add("heart")

	node = rootTrie.Find("heart")

	node.Set("test", "alpha")
	_, exists = node.Get("tests")

	if exists == true {
		panic("a problem")
	}
}
