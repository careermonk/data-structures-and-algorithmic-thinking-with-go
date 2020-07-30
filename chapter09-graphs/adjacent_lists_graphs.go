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
	"errors"
	"fmt"
)

type GraphType string

const (
	DIRECTED   GraphType = "DIRECTED"
	UNDIRECTED GraphType = "UNDIRECTED"
)

type Graph interface {
	Init()

	AddEdge(vertexOne int, vertexTwo int) error

	AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) error

	RemoveEdge(vertexOne int, vertexTwo int) error

	HasEdge(vertexOne int, vertexTwo int) bool

	GetGraphType() GraphType

	GetAdjacentNodesForVertex(vertex int) map[int]bool

	GetWeightOfEdge(vertexOne int, vertexTwo int) (int, error)

	GetNumberOfVertices() int

	GetNumberOfEdges() int

	GetIndegreeForVertex(vertex int) int
}

type AdjacencyList struct {
	Vertices  int
	Edges     int
	GraphType GraphType
	AdjList   []*Node
}

type Node struct {
	Next   *Node
	Weight int
	Key    int
}

// Recursive method to add node to last available slot
func (node Node) AddNode(value int) *Node {
	n := node.Next
	if n == nil {
		newnode := &Node{Next: &Node{}, Key: value}
		return newnode
	}
	nd := n.AddNode(value)
	node.Next = nd
	return &node
}

// Recursive method to append with weight
func (node Node) AddNodeWithWeight(value int, weight int) *Node {
	n := node.Next
	if n == nil {
		newnode := &Node{Next: &Node{}, Key: value, Weight: weight}
		return newnode
	}
	nd := n.AddNodeWithWeight(value, weight)
	node.Next = nd
	return &node
}

func (node Node) FindNextNode(key int) (*Node, error) {
	n := node
	if n == (Node{}) {
		return &Node{}, errors.New("Node not found")
	}
	if n.Key == key {
		return &n, nil
	}
	nd := n.Next
	return nd.FindNextNode(key)
}

func (G *AdjacencyList) Init() {
	G.AdjList = make([]*Node, G.Vertices)
	G.Edges = 0
	for i := 0; i < G.Vertices; i++ {
		G.AdjList[i] = &Node{}
	}
}

func (G *AdjacencyList) AddEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}
	node := G.AdjList[vertexOne].AddNode(vertexTwo)
	G.AdjList[vertexOne] = node
	if G.GraphType == UNDIRECTED {
		node := G.AdjList[vertexTwo].AddNode(vertexOne)
		G.AdjList[vertexTwo] = node
	}
	G.Edges++
	return nil
}

func (G *AdjacencyList) AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}
	node := G.AdjList[vertexOne].AddNodeWithWeight(vertexTwo, weight)
	G.AdjList[vertexOne] = node
	if G.GraphType == UNDIRECTED {
		node := G.AdjList[vertexTwo].AddNodeWithWeight(vertexOne, weight)
		G.AdjList[vertexTwo] = node
	}
	G.Edges++
	return nil
}

func (G *AdjacencyList) RemoveEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}
	nodeAdj := G.AdjList[vertexOne]
	if nodeAdj == (&Node{}) {
		return errors.New("Node not found")
	}
	nextNode := nodeAdj
	newNodes := Node{}
	for nextNode != (&Node{}) && nextNode != nil {
		if nextNode.Key != vertexTwo {
			newNodes.Next = nextNode
			newNodes.Key = nextNode.Key
			newNodes.Weight = nextNode.Weight
			nextNode = nextNode.Next
		} else {
			newNodes.Next = nextNode.Next
			newNodes.Key = nextNode.Next.Key
			newNodes.Weight = nextNode.Next.Weight
			G.AdjList[vertexOne] = &newNodes
			G.Edges--
			return nil
		}
	}
	G.Edges--
	return nil
}

func (G *AdjacencyList) HasEdge(vertexOne int, vertexTwo int) bool {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return false
	}
	nodeAdj := G.AdjList[vertexOne]
	if nodeAdj == (&Node{}) {
		return false
	}
	node, _ := nodeAdj.FindNextNode(vertexTwo)
	if node != nil && node.Key == vertexTwo {
		return true
	}
	return false
}

func (G *AdjacencyList) GetGraphType() GraphType {
	return G.GraphType
}

func (G *AdjacencyList) GetAdjacentNodesForVertex(vertex int) map[int]bool {
	if vertex >= G.Vertices || vertex < 0 {
		return map[int]bool{}
	}
	nodeAdj := G.AdjList[vertex]
	nextNode := nodeAdj
	nodes := map[int]bool{}
	for nextNode != (&Node{}) && nextNode != nil {
		nodes[nextNode.Key] = true
		nextNode = nextNode.Next
	}
	return nodes
}

func (G *AdjacencyList) GetWeightOfEdge(vertexOne int, vertexTwo int) (int, error) {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return 0, errors.New("Error getting weight for vertex")
	}
	nodeAdj := G.AdjList[vertexOne]
	if nodeAdj == (&Node{}) {
		return 0, errors.New("Error getting weight for vertex")
	}
	node, _ := nodeAdj.FindNextNode(vertexTwo)
	if node != nil && node.Key == vertexTwo {
		return nodeAdj.Weight, nil
	}
	return 0, errors.New("Error getting weight for vertex")
}

func (G *AdjacencyList) GetNumberOfVertices() int {
	return G.Vertices
}

func (G *AdjacencyList) GetNumberOfEdges() int {
	return G.Edges
}

func (G *AdjacencyList) GetIndegreeForVertex(vertex int) int {
	if vertex >= G.Vertices || vertex < 0 {
		return 0
	}
	nodeAdj := G.AdjList[vertex]
	nextNode := nodeAdj.Next
	length := 0
	for nextNode != (&Node{}) && nextNode != nil {
		length += 1
		nextNode = nextNode.Next
	}
	return length
}

func main() {
	TestAdjacencyList_AddEdgeDirectedGreaterThanVertex()
	TestAdjacencyList_AddEdgeDirected()
	TestAdjacencyList_AddEdgeUndirected()
	TestAdjacencyList_AddEdgeWeightDirectedGreaterThanVertex()
	TestAdjacencyList_AddEdgeWithWeightDirected()
	TestAdjacencyList_AddEdgeWithWeightUnDirected()
	TestAdjacencyList_RemoveEdgeDirected()
	TestAdjacencyList_RemoveEdgeUnDirected()
	TestAdjacencyList_RemoveEdgeDirected()
	TestAdjacencyList_RemoveEdgeUnDirected()
	TestAdjacencyList_HasEdgeDirected()
	TestAdjacencyList_HasEdgeUndirected()
	TestAdjacencyList_GetGraphType()
	TestAdjacencyList_GetAdjacentVerticesNodesForVertexUndirected()
	TestAdjacencyList_GetWeightOfEdgeDirected()
	TestAdjacencyList_GetWeightOfEdgeUndirected()
	TestAdjacencyList_GetNumberOfVertices()
	TestAdjacencyList_GetNumberOfEdges()
	TestAdjacencyList_GetIndegreeForVertex()
}

// Tests
var testAdjListDirected = &AdjacencyList{4, 0, DIRECTED, nil}
var testAdjListUnDirected = &AdjacencyList{4, 0, UNDIRECTED, nil}

func TestAdjacencyList_AddEdgeDirectedGreaterThanVertex() {
	testAdjListDirected.Init()
	err := testAdjListDirected.AddEdge(10, 2)
	if err == nil {
		fmt.Printf("Index was out of bounds it should have failed")
	}
}

func TestAdjacencyList_AddEdgeDirected() {
	testAdjListDirected.Init()
	testAdjListDirected.AddEdge(2, 1)
	err := testAdjListDirected.AddEdge(2, 3)
	if err != nil {
		fmt.Printf("Error adding edge")
	}
	if testAdjListDirected.AdjList[2].Key != 1 {
		fmt.Printf("Data not found at index")
	}
	if testAdjListDirected.AdjList[2].Next.Key != 3 {
		fmt.Printf("Data not found at index")
	}
}

func TestAdjacencyList_AddEdgeUndirected() {
	testAdjListUnDirected.Init()
	testAdjListUnDirected.AddEdge(2, 1)
	testAdjListUnDirected.AddEdge(2, 3)
	err := testAdjListUnDirected.AddEdge(2, 0)
	if err != nil {
		fmt.Printf("Error adding edge " + err.Error())
	}
	if testAdjListUnDirected.AdjList[2].Key != 1 {
		fmt.Printf("Data not found at index")
	}
	if testAdjListUnDirected.AdjList[2].Next.Key != 3 {
		fmt.Printf("Data not found at index")
	}
	if testAdjListUnDirected.AdjList[1].Key != 2 {
		fmt.Printf("Data not found at index")
	}
	if testAdjListUnDirected.AdjList[3].Key != 2 {
		fmt.Printf("Data not found at index")
	}
}

func TestAdjacencyList_AddEdgeWeightDirectedGreaterThanVertex() {
	testAdjListUnDirected.Init()
	err := testAdjListUnDirected.AddEdgeWithWeight(10, 7, -3)
	if err == nil {
		fmt.Printf("Index was out of bounds it should have failed")
	}
}

func TestAdjacencyList_AddEdgeWithWeightDirected() {
	testAdjListDirected.Init()
	testAdjListDirected.AddEdgeWithWeight(2, 1, -3)
	err := testAdjListDirected.AddEdgeWithWeight(2, 3, 6)
	if err != nil {
		fmt.Printf("Error adding edge")
	}
	if testAdjListDirected.AdjList[2].Weight != -3 {
		fmt.Printf("Data not found at index")
	}
	if testAdjListDirected.AdjList[2].Next.Weight != 6 {
		fmt.Printf("Data not found at index, %d", testAdjListUnDirected.AdjList[2].Next.Weight)
	}
}

func TestAdjacencyList_AddEdgeWithWeightUnDirected() {
	testAdjListUnDirected.Init()
	testAdjListUnDirected.AddEdgeWithWeight(2, 1, -3)
	err := testAdjListUnDirected.AddEdgeWithWeight(2, 3, 6)
	if err != nil {
		fmt.Printf("Error adding edge")
	}
	if testAdjListUnDirected.AdjList[2].Weight != -3 {
		fmt.Printf("Data not found at index")
	}
	if testAdjListUnDirected.AdjList[1].Weight != -3 {
		fmt.Printf("Data not found at index")
	}
	if testAdjListUnDirected.AdjList[2].Next.Weight != 6 {
		fmt.Printf("Data not found at index, %d", testAdjListUnDirected.AdjList[2].Next.Weight)
	}
	if testAdjListUnDirected.AdjList[3].Weight != 6 {
		fmt.Printf("Data not found at index, %d", testAdjListUnDirected.AdjList[2].Next.Weight)
	}
}

func TestAdjacencyList_RemoveEdgeDirected() {
	testAdjListDirected.Init()
	testAdjListDirected.AddEdge(2, 1)
	testAdjListDirected.AddEdge(2, 3)
	testAdjListDirected.AddEdge(2, 0)
	err := testAdjListDirected.RemoveEdge(2, 3)
	if err != nil {
		fmt.Printf("Error removing edge")
	}
	if testAdjListDirected.AdjList[2].Key != 1 && testAdjListDirected.AdjList[2].Next.Key != 0 && testAdjListDirected.AdjList[2].Next.Next != (&Node{}) {
		fmt.Printf("Data not removed at index, data is %d instead of 0, %v", testAdjListDirected.AdjList[2].Next.Key, testAdjListDirected.AdjList)
	}
}

func TestAdjacencyList_RemoveEdgeUnDirected() {
	testAdjListUnDirected.Init()
	testAdjListUnDirected.AddEdge(2, 1)
	testAdjListUnDirected.AddEdge(2, 3)
	testAdjListUnDirected.AddEdge(2, 0)
	err := testAdjListUnDirected.RemoveEdge(2, 3)
	if err != nil {
		fmt.Printf("Error removing edge")
	}
	if testAdjListUnDirected.HasEdge(2, 3) {
		fmt.Printf("There's still a relationship between nodes")
	}
	if testAdjListUnDirected.AdjList[2].Key != 1 && testAdjListUnDirected.AdjList[2].Next.Key != 0 && testAdjListUnDirected.AdjList[2].Next.Next != (&Node{}) {
		fmt.Printf("Data not removed at index, data is %d instead of 0, %v", testAdjListUnDirected.AdjList[2].Next.Key, testAdjListUnDirected.AdjList)
	}
}

func TestAdjacencyList_HasEdgeDirected() {
	testAdjListDirected.Init()
	testAdjListDirected.AddEdge(2, 1)
	testAdjListDirected.AddEdge(2, 3)
	testAdjListDirected.AddEdge(2, 0)
	if !testAdjListDirected.HasEdge(2, 0) {
		fmt.Printf("No relationship, when there should be one")
	}
	if testAdjListDirected.HasEdge(1, 2) {
		fmt.Printf("Relationship, when there shouldn't be one")
	}
}

func TestAdjacencyList_HasEdgeUndirected() {
	testAdjListUnDirected.Init()
	testAdjListUnDirected.AddEdge(2, 1)
	testAdjListUnDirected.AddEdge(2, 3)
	testAdjListUnDirected.AddEdge(2, 0)
	if !testAdjListUnDirected.HasEdge(2, 0) {
		fmt.Printf("No relationship, when there should be one")
	}
	if !testAdjListUnDirected.HasEdge(3, 2) {
		fmt.Printf("No relationship, when there should be one")
	}
}

func TestAdjacencyList_GetGraphType() {
	testAdjListDirected.Init()
	if testAdjListDirected.GraphType != testAdjListDirected.GetGraphType() {
		fmt.Printf("GraphType not matching")
	}
}

func TestAdjacencyList_GetAdjacentVerticesNodesForVertexUndirected() {
	testAdjListUnDirected.Init()
	testAdjListUnDirected.AddEdge(2, 1)
	testAdjListUnDirected.AddEdge(2, 0)
	testAdjListUnDirected.AddEdge(1, 3)
	testAdjListUnDirected.AddEdge(2, 2)
	testAdjListUnDirected.AddEdge(2, 3)
	nodes := testAdjListUnDirected.GetAdjacentNodesForVertex(2)
	if len(nodes) != 4 {
		fmt.Printf("Nodes size not matching. Size %d instead of %d. %v", len(nodes), 4, nodes)
	}
}

func TestAdjacencyList_GetWeightOfEdgeDirected() {
	testAdjListDirected.Init()
	testAdjListDirected.AddEdgeWithWeight(2, 1, -3)
	weight, _ := testAdjListDirected.GetWeightOfEdge(2, 1)
	if weight != -3 {
		fmt.Printf("Data not found at index")
	}
}

func TestAdjacencyList_GetWeightOfEdgeUndirected() {
	testAdjListUnDirected.Init()
	testAdjListUnDirected.AddEdgeWithWeight(2, 1, -3)
	weight, _ := testAdjListUnDirected.GetWeightOfEdge(2, 1)
	if weight != -3 {
		fmt.Printf("Data not found at index")
	}
}

func TestAdjacencyList_GetNumberOfVertices() {
	testAdjListUnDirected.Init()
	if testAdjListUnDirected.GetNumberOfVertices() != testAdjListUnDirected.Vertices {
		fmt.Printf("Vertex count doesn't match")
	}
}

func TestAdjacencyList_GetNumberOfEdges() {
	testAdjListDirected.Init()
	testAdjListDirected.AddEdgeWithWeight(2, 1, -3)
	if testAdjListDirected.GetNumberOfEdges() != 1 {
		fmt.Printf("Edges count doesn't match")
	}
}

func TestAdjacencyList_GetIndegreeForVertex() {
	testAdjListDirected.Init()
	testAdjListDirected.AddEdge(2, 1)
	testAdjListDirected.AddEdge(2, 0)
	testAdjListDirected.AddEdge(1, 1)
	testAdjListDirected.AddEdge(2, 3)
	testAdjListDirected.AddEdge(0, 3)
	if testAdjListDirected.GetIndegreeForVertex(2) != 3 {
		fmt.Printf("Nodes size not matching. Expected %d but is %d ", 3, testAdjListDirected.GetIndegreeForVertex(2))
	}
}
