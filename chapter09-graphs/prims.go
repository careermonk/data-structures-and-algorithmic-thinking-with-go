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
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type minHeapNode interface {
	lessThan(minHeapNode) bool
}

type MinHeap struct {
	nodes []minHeapNode
}

func NewMinHeap(nodes []minHeapNode) *MinHeap {
	minHeap := &MinHeap{
		nodes: nodes,
	}

	if len(minHeap.nodes) <= 1 {
		return minHeap
	}

	// floyds algorithm is used to guarantee the heap property when a new
	// binary heap is created. This is down by starting at the bottom of
	// the heap and sifing nodes so the heap property is fulfilled
	for i := len(minHeap.nodes) / 2; i >= 0; i-- {
		minHeap.percolateDown(i)
	}

	return minHeap
}

func (mh *MinHeap) Insert(node minHeapNode) {
	mh.nodes = append(mh.nodes, node)
	mh.percolateUp(len(mh.nodes) - 1)
}

func (mh *MinHeap) percolateUp(index int) {
	// ensure that the index is in range
	if index == 0 || index >= len(mh.nodes) {
		return
	}

	parentIndex := (index - 1) / 2
	// if the child node is less than the parent, then swap the nodes in the backing array
	if mh.nodes[index].lessThan(mh.nodes[parentIndex]) {
		temp := mh.nodes[parentIndex]
		mh.nodes[parentIndex] = mh.nodes[index]
		mh.nodes[index] = temp
		mh.percolateUp(parentIndex)
	}
}

func (mh *MinHeap) percolateDown(index int) {
	// there isn't a single child, so stop sifting down
	if index*2+1 >= len(mh.nodes) {
		return
	}

	childIndex := index*2 + 1
	childNode := mh.nodes[index*2+1]

	// see if the second child node is less than the first, and if so swap it out
	if index*2+2 < len(mh.nodes) && mh.nodes[index*2+2].lessThan(childNode) {
		childNode = mh.nodes[index*2+2]
		childIndex = index*2 + 2
	}

	// if the childNode is less than the parent, then we need to swap them out
	if childNode.lessThan(mh.nodes[index]) {
		temp := mh.nodes[index]
		mh.nodes[index] = mh.nodes[childIndex]
		mh.nodes[childIndex] = temp
		mh.percolateDown(childIndex)
	}
}

func (mh *MinHeap) Pop() (minHeapNode, error) {
	if len(mh.nodes) == 0 {
		return nil, fmt.Errorf("MinHeap is empty")
	}

	node := mh.nodes[0]
	mh.nodes[0] = mh.nodes[len(mh.nodes)-1]
	mh.nodes = mh.nodes[:len(mh.nodes)-1]
	mh.percolateDown(0)

	return node, nil
}

func (mh MinHeap) Fetch() (minHeapNode, error) {
	if len(mh.nodes) == 0 {
		return nil, fmt.Errorf("MinHeap is empty")
	}

	return mh.nodes[0], nil
}

func (mh MinHeap) Size() int {
	return len(mh.nodes)
}

type Node interface{}

type ObjectNode struct {
	name string
}

func (o ObjectNode) String() string {
	return fmt.Sprintf("%s", o.name)
}

type edge struct {
	nodes  [2]Node
	weight int
}

func (e edge) String() string {
	return fmt.Sprintf("%s<-%d->%s", e.nodes[0], e.weight, e.nodes[1])
}

func (e edge) lessThan(mh minHeapNode) bool {
	edge, ok := mh.(*edge)
	if !ok {
		log.Fatal("Programming error. Invalid type assertion of minHeapNode type to *edge type")
	}

	return e.weight < edge.weight
}

// A graph is an object which represents an undirected graph which has many
// edges connecting different nodes.
type Graph struct {
	// edges are stored as a map of minHeaps, for each
	edges map[Node]*MinHeap
}

func NewGraph() *Graph {
	return &Graph{
		edges: make(map[Node]*MinHeap, 0),
	}
}

func (g *Graph) Insert(nodeA Node, nodeB Node, weight int) {
	// look up and insert the first Node; creating the edge in the process
	insert := func(node Node, edge *edge) {
		nodeEdges, ok := g.edges[node]
		if !ok {
			g.edges[node] = NewMinHeap([]minHeapNode{})
			nodeEdges = g.edges[node]
		}

		nodeEdges.Insert(edge)
	}

	edge := &edge{
		nodes:  [2]Node{nodeA, nodeB},
		weight: weight,
	}
	insert(nodeA, edge)
	insert(nodeB, edge)
}

func (g Graph) Get(node Node) (*MinHeap, error) {
	nodeEdges, ok := g.edges[node]
	if !ok {
		return nil, fmt.Errorf("Node not found")
	}

	return nodeEdges, nil
}

func (g *Graph) IterNodes(cb func(Node) bool) {
	for node, _ := range g.edges {
		if shouldContinue := cb(node); !shouldContinue {
			break
		}
	}
}

func (g *Graph) IterEdges(cb func(*edge) bool) {
	edges := make(map[*edge]interface{})
	for _, nodeEdges := range g.edges {
		poppedEdges := make([]*edge, 0)
		for {
			heapObj, err := nodeEdges.Pop()
			if err != nil {
				break
			}

			edge := heapObj.(*edge)
			poppedEdges = append(poppedEdges, edge)
			if _, ok := edges[edge]; !ok {
				edges[edge] = struct{}{}
				cb(edge)
			}
		}

		for _, edge := range poppedEdges {
			nodeEdges.Insert(edge)
		}
	}
}

// A tree type is a type which represents a set of vertices which are all only
// connected in one manner. This particular tree type represents a complete
// tree which is non cyclic and where each node only has a maximum of two
// edges.
type Tree struct {
	edges map[Node]map[Node]*edge
}

func NewTree() *Tree {
	return &Tree{
		edges: make(map[Node]map[Node]*edge, 0),
	}
}

func (t *Tree) Insert(nodeA, nodeB Node, nodeEdge *edge) error {
	if _, nodeAExists := t.edges[nodeA]; nodeAExists {
		if _, nodeBExists := t.edges[nodeB]; nodeBExists {
			return fmt.Errorf("both nodes already exist in the tree")
		}
	}

	insert := func(node Node, connectedNode Node, e *edge) (func(), error) {
		nodeEdges, ok := t.edges[node]
		if !ok {
			t.edges[node] = make(map[Node]*edge, 2)
			nodeEdges = t.edges[node]
		}

		if _, alreadyExists := nodeEdges[connectedNode]; alreadyExists {
			return nil, fmt.Errorf("Can not insert same node into the tree twice")
		}

		if len(nodeEdges) > 1 {
			return nil, fmt.Errorf("Each node can only have two edges in tree")
		}

		commit := func() {
			nodeEdges[connectedNode] = e
		}

		return commit, nil
	}

	commitA, err := insert(nodeA, nodeB, nodeEdge)
	if err != nil {
		return err
	}

	commitB, err := insert(nodeB, nodeA, nodeEdge)
	if err != nil {
		return err
	}

	// if both insertions are permitted, then go ahead and commit each
	// change to the the tree. This is done to simplify the implementation
	// so as to allow us to verify that this is a valid operation for
	// _both_ nodes before we commit either.
	commitA()
	commitB()

	return nil
}

func (t Tree) IterEdges(cb func(edge *edge) bool) {
	currentNode := func() Node {
		for node, edges := range t.edges {
			if len(edges) == 1 {
				return node
			}
		}

		return nil
	}()

	if currentNode == nil {
		return
	}

	otherNode := func(node Node, edge *edge) Node {
		if node == edge.nodes[0] {
			return edge.nodes[1]
		}

		return edge.nodes[0]
	}

	traversedEdges := make(map[*edge]interface{}, 0)
	for {
		edges := t.edges[currentNode]
		found := false
		for _, edge := range edges {
			if _, ok := traversedEdges[edge]; ok {
				continue
			}
			currentNode = otherNode(currentNode, edge)
			traversedEdges[edge] = struct{}{}
			cb(edge)
			found = true
		}

		if !found {
			break
		}
	}
}

func Prims(g *Graph) (*Tree, error) {
	// this is an implementation of Prim's algorithm, which iterates
	// through the graph and builds a tree, node by node by choosing the
	// next node with the smallest edge that doesn't already exist in the tree.

	// requiredNodes returns a map of all of the nodes that are required to build the MST.
	requiredNodes := func() map[Node]interface{} {
		nodes := make(map[Node]interface{}, 0)
		iter := func(node Node) bool {
			nodes[node] = struct{}{}
			return true
		}

		g.IterNodes(iter)
		return nodes
	}()

	// otherNode returns the other node from an Edge; this is useful so we can parse
	// the underlying datatype and compose things together as we have done thus far.
	otherEdgeNode := func(node Node, e *edge) Node {
		if e.nodes[0] == node {
			return e.nodes[1]
		}

		return e.nodes[0]
	}

	tree := NewTree()
	nodes := make([]Node, 0, len(requiredNodes))

	current := func() Node {
		for key, _ := range requiredNodes {
			return key
		}
		// should never be called unless requiredNodes is empty
		return nil
	}()

	for {
		if len(nodes) == len(requiredNodes)-1 {
			break
		}

		nodeEdges, err := g.Get(current)
		if err != nil {
			return nil, err
		}

		poppedEdges := make([]*edge, 0, nodeEdges.Size())

		for {
			minHeapItem, err := nodeEdges.Pop()
			// if there is nothing to pop from the tree, then that means the algorithm was too greedy and is failing!
			if err != nil {
				return nil, fmt.Errorf("Unable to find a Node to continue building the tree")
			}

			minWeightEdge, ok := minHeapItem.(*edge)
			if !ok {
				log.Fatal("Programming error. Type assertion was attempted which is impossible")
			}

			poppedEdges = append(poppedEdges, minWeightEdge)
			newNode := otherEdgeNode(current, minWeightEdge)

			// if the insertion was successful, then we updated the tree successfully and can exit
			if err := tree.Insert(current, newNode, minWeightEdge); err == nil {
				// update the current iterator to the new node
				current = newNode
				nodes = append(nodes, newNode)
				break
			}
		}

		// we put all of the popped edges back into the heap, so this doesn't arbitrarily change the underlying
		// heap. Practically speaking, this is a bit of a code smell so this needs a little refactoring.
		for _, poppedEdge := range poppedEdges {
			nodeEdges.Insert(poppedEdge)
		}
	}
	return tree, nil
}

// Test code
func main() {
	TestBuildMinHeap()
	TestGraphInsertion()
	TestTree()
	TestMSTImplementation()
}

func init() {
	rand.Seed(time.Now().Unix())
}

func randomList() []int {
	// random list between 500 - 1000 items
	len := int(rand.Float32()*500) + 500
	len = 100
	list := make([]int, 0, len)
	for i := 0; i < len; i++ {
		list = append(list, rand.Intn(100))
	}

	return list
}

type minHeapTestNode struct {
	value int
}

func (mh minHeapTestNode) lessThan(node minHeapNode) bool {
	testNode := node.(minHeapTestNode)
	return mh.value < testNode.value
}

func TestBuildMinHeap() {
	for i := 0; i < 1; i++ {
		nodes := make([]minHeapNode, 0)
		for _, val := range randomList() {
			nodes = append(nodes, minHeapTestNode{val})
		}

		minHeap := NewMinHeap(nodes)
		processed := 1
		lastNode, err := minHeap.Pop()
		if err != nil {
			fmt.Println("Did not build heap properly")
		}

		for {
			node, err := minHeap.Pop()
			if err != nil {
				break
			}

			if node.lessThan(lastNode) {
				fmt.Println("MinHeap does not fulfill the MinHeap property")
			}

			if err != nil && processed != len(nodes) {
				fmt.Println("Did not pop all nodes off properly")
				break
			}

			lastNode = node
		}
	}
}

func TestGraphInsertion() {
	mapGraph := make(map[[2]string]int, 0)
	mapGraph[[2]string{"a", "b"}] = 5
	mapGraph[[2]string{"a", "c"}] = 8
	mapGraph[[2]string{"a", "d"}] = 1
	mapGraph[[2]string{"a", "f"}] = 5

	mapGraph[[2]string{"b", "g"}] = 9
	mapGraph[[2]string{"b", "h"}] = 1
	mapGraph[[2]string{"b", "c"}] = 5

	mapGraph[[2]string{"c", "d"}] = 9
	mapGraph[[2]string{"c", "e"}] = 1
	mapGraph[[2]string{"c", "f"}] = 5

	mapGraph[[2]string{"d", "h"}] = 1
	mapGraph[[2]string{"d", "e"}] = 9

	mapGraph[[2]string{"e", "f"}] = 9
	mapGraph[[2]string{"e", "h"}] = 2

	graph := NewGraph()
	for nodes, weight := range mapGraph {
		graph.Insert(nodes[0], nodes[1], weight)
	}

	nodes := make(map[string]interface{}, 0)
	iter := func(n Node) bool {
		node, ok := n.(string)
		if !ok {
			fmt.Println("Programming error, the node was not found correctly")
		}

		nodes[node] = struct{}{}
		return true
	}

	graph.IterNodes(iter)
	if len(nodes) != 8 {
		fmt.Println("Graph did not store and house nodes correctly")
	}
}

func TestTree() {
	testCases := []struct {
		nodes     [2]string
		weight    int
		shouldErr bool
	}{
		{[2]string{"a", "b"}, 2, false},
		{[2]string{"b", "c"}, 4, false},
		{[2]string{"c", "d"}, 2, false},
		{[2]string{"c", "d"}, 2, true},
		{[2]string{"d", "e"}, 4, false},
		{[2]string{"e", "f"}, 2, false},
		{[2]string{"c", "e"}, 2, true},
	}

	expectedWeight := 0
	tree := NewTree()
	nodes := make(map[string]bool, 0)

	for _, testCase := range testCases {
		err := tree.Insert(
			testCase.nodes[0],
			testCase.nodes[1],
			&edge{
				[2]Node{testCase.nodes[0], testCase.nodes[1]}, testCase.weight,
			})

		if err != nil && !testCase.shouldErr {
			fmt.Println("Tree insertion failed when it shouldn't have")
		}

		if err == nil && testCase.shouldErr {
			fmt.Println("Tree insertion did not fail when expected")
		}

		if !testCase.shouldErr {
			expectedWeight += testCase.weight
			nodes[testCase.nodes[0]] = false
			nodes[testCase.nodes[1]] = false
		}
	}

	iter := func(edge *edge) bool {
		nodes[edge.nodes[0].(string)] = true
		nodes[edge.nodes[1].(string)] = true
		expectedWeight -= edge.weight
		return true
	}
	tree.IterEdges(iter)

	if expectedWeight != 0 {
		fmt.Println("Tree iteration did not return correct overall weight")
	}

	for _, value := range nodes {
		if value != true {
			fmt.Println("Some nodes were not stored properly")
		}
	}
}

func buildGraph(totalNodes int) *Graph {
	nodes := make([]string, 0, totalNodes)
	graph := NewGraph()

	for i := 0; i < totalNodes; i++ {
		// get ascii encoded version of number so we can encode it
		node := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(i)))
		node = strings.Replace(node, "=", "", -1)

		nodes = append(nodes, node)
	}

	for index, node := range nodes {
		for siblingIndex, siblingNode := range nodes {
			// generate a score between 10 and 1000
			if siblingIndex == index {
				continue
			}
			score := rand.Intn(1000-10) + 10
			graph.Insert(node, siblingNode, score)
		}
	}

	return graph
}

func TestMSTImplementation() {
	graph := buildGraph(100)
	_, err := Prims(graph)
	if err != nil {
		fmt.Printf("%s", err)
	}
}
