// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

package main

import "fmt"

type Element struct {
	parent *Element    // Parent element
	rank   int         // Rank (approximate depth) of the subtree with this element as root
	Data   interface{} // Arbitrary user-provided payload
}

// MakeSet creates a singleton set and returns its sole element.
func MakeSet(Data interface{}) *Element {
	s := &Element{}
	s.parent = s
	s.Data = Data
	return s
}

func Find(e *Element) *Element {
	for e.parent != e {
		e.parent = e.parent.parent
		e = e.parent
	}
	return e
}

func FindR(e *Element) *Element {
	if e.parent == e {
		return e
	} else {
		return FindR(e.parent)
	}
}

func Union(e1, e2 *Element) {
	e1.parent = e2
}

func main() {
	aSet := MakeSet("a")
	bSet := MakeSet("b")
	oneSet := MakeSet(1)
	twoSet := MakeSet(2)
	Union(aSet, bSet)
	Union(oneSet, twoSet)
	result := Find(aSet)
	fmt.Println(result.Data)
	result = Find(bSet)
	fmt.Println(result.Data)
	result = Find(oneSet)
	fmt.Println(result.Data)
	result = Find(twoSet)
	fmt.Println(result.Data)
}
