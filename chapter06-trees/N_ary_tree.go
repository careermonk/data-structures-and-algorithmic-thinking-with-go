package main

import (
	"fmt"
)

// Total is a fun way to total how many nodes we have
var total = 1

// N-aray
const N = 3

// NaryTreeNode  is a super simple NaryTreeNode  struct that will form the tree
type NaryTreeNode struct {
	parent   *NaryTreeNode
	children []*NaryTreeNode
	depth    int
	data     int
}

// BuildNaryTree will construct the tree for walking.
// total = (N * childrenChildren) + rootChildren + 1
func BuildNaryTree() *NaryTreeNode {
	var root NaryTreeNode
	root.addChildren()

	for _, child := range root.children {
		child.addChildren()
	}
	return &root
}

// addChildren is a convenience to add an arbitrary number of children
func (n *NaryTreeNode) addChildren() {
	for i := 0; i < N; i++ {
		newChild := &NaryTreeNode{
			parent: n,
			depth:  n.depth + 1,
		}
		n.children = append(n.children, newChild)
	}
}

// addChild is a convenience to add an a children
func (n *NaryTreeNode) addChild() {
	newChild := &NaryTreeNode{
		parent: n,
		depth:  n.depth + 1,
	}
	n.children = append(n.children, newChild)
}

// walk is a recursive function that calls itself for every child
func (n *NaryTreeNode) walk() {
	n.visit()
	for _, child := range n.children {
		child.walk()
	}
}

// visit will get called on every NaryTreeNode  in the tree.
func (n *NaryTreeNode) visit() {
	d := "└"
	for i := 0; i <= n.depth; i++ {
		d = d + "───"
	}
	fmt.Printf("%s Visiting NaryTreeNode  with address %p and parent %p Total (%d)\n", d, n, n.parent, total)
	total = total + 1
}

// main will build and walk the tree
func main() {
	fmt.Println("Building N-ary Tree")
	root := BuildNaryTree()
	fmt.Println("Walking N-ary Tree")
	root.walk()

}
