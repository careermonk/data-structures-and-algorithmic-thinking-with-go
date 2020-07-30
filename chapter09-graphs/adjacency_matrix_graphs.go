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

type AdjacencyMatrix struct {
	Vertices  int
	Edges     int
	GraphType GraphType
	AdjMatrix [][]int
}

func (G *AdjacencyMatrix) Init() {
	G.AdjMatrix = make([][]int, G.Vertices)
	G.Edges = 0
	for i := 0; i < G.Vertices; i++ {
		G.AdjMatrix[i] = make([]int, G.Vertices)
	}
}

func (G *AdjacencyMatrix) AddEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}

	G.AdjMatrix[vertexOne][vertexTwo] = 1
	G.Edges++
	if G.GraphType == UNDIRECTED {
		G.AdjMatrix[vertexTwo][vertexOne] = 1
		G.Edges++
	}
	return nil
}

func (G *AdjacencyMatrix) AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}

	G.AdjMatrix[vertexOne][vertexTwo] = weight
	G.Edges++
	if G.GraphType == UNDIRECTED {
		G.AdjMatrix[vertexTwo][vertexOne] = weight
		G.Edges++
	}
	return nil
}

func (G *AdjacencyMatrix) RemoveEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}

	G.AdjMatrix[vertexOne][vertexTwo] = 0
	G.Edges--
	if G.GraphType == UNDIRECTED {
		G.AdjMatrix[vertexTwo][vertexOne] = 0
		G.Edges--
	}
	return nil
}

func (G *AdjacencyMatrix) HasEdge(vertexOne int, vertexTwo int) bool {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return false
	}
	return G.AdjMatrix[vertexOne][vertexTwo] != 0
}

func (G *AdjacencyMatrix) GetGraphType() GraphType {
	return G.GraphType
}

func (G *AdjacencyMatrix) GetAdjacentNodesForVertex(vertex int) map[int]bool {
	adjacencyMatrixVertices := map[int]bool{}
	if vertex >= G.Vertices || vertex < 0 {
		return adjacencyMatrixVertices
	}
	for i := 0; i < G.Vertices; i++ {
		if G.AdjMatrix[vertex][i] != 0 {
			adjacencyMatrixVertices[i] = (G.AdjMatrix[vertex][i] != 0)
		}
	}
	return adjacencyMatrixVertices
}

func (G *AdjacencyMatrix) GetWeightOfEdge(vertexOne int, vertexTwo int) (int, error) {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return 0, errors.New("Error getting weight for vertex")
	}
	return G.AdjMatrix[vertexOne][vertexTwo], nil
}

func (G *AdjacencyMatrix) GetNumberOfVertices() int {
	return G.Vertices
}

func (G *AdjacencyMatrix) GetNumberOfEdges() int {
	return G.Edges
}

func (G *AdjacencyMatrix) GetIndegreeForVertex(vertex int) int {
	indegree := 0
	adjacentNodes := G.GetAdjacentNodesForVertex(vertex)
	for key := range adjacentNodes {
		if adjacentNodes[key] {
			indegree++
		}
	}
	return indegree
}

// Tests

func main() {
	TestAdjacencyMatrix_AddEdgeDirectedGreaterThanVertex()
	TestAdjacencyMatrix_AddEdgeDirected()
	TestAdjacencyMatrix_AddEdgeUndirected()
	TestAdjacencyMatrix_AddEdgeWeightDirectedGreaterThanVertex()
	TestAdjacencyMatrix_AddEdgeWithWeightDirected()
	TestAdjacencyMatrix_AddEdgeWithWeightUnDirected()
	TestAdjacencyMatrix_RemoveEdgeDirected()
	TestAdjacencyMatrix_RemoveEdgeUnDirected()
	TestAdjacencyMatrix_RemoveEdgeDirected()
	TestAdjacencyMatrix_RemoveEdgeUnDirected()
	TestAdjacencyMatrix_HasEdgeDirected()
	TestAdjacencyMatrix_HasEdgeUndirected()
	TestAdjacencyMatrix_GetGraphType()
	TestAdjacencyMatrix_GetAdjacentNodesForVertexDirected()
	TestAdjacencyMatrix_GetWeightOfEdgeDirected()
	TestAdjacencyMatrix_GetWeightOfEdgeUndirected()
	TestAdjacencyMatrix_GetNumberOfVertices()
	TestAdjacencyMatrix_GetNumberOfEdges()
	TestAdjacencyMatrix_GetIndegreeForVertex()
}

var testAdjMatrixDirected = &AdjacencyMatrix{4, 0, DIRECTED, nil}
var testAdjMatrixUnDirected = &AdjacencyMatrix{4, 0, UNDIRECTED, nil}

func TestAdjacencyMatrix_AddEdgeDirectedGreaterThanVertex() {
	testAdjMatrixDirected.Init()
	err := testAdjMatrixDirected.AddEdge(10, 2)
	if err == nil {
		fmt.Printf("Index was out of bounds it should have failed")
	}
}

func TestAdjacencyMatrix_AddEdgeDirected() {
	testAdjMatrixDirected.Init()
	err := testAdjMatrixDirected.AddEdge(2, 1)
	if err != nil {
		fmt.Printf("Error adding edge")
	}
	if testAdjMatrixDirected.AdjMatrix[2][1] != 1 {
		fmt.Printf("Data not found at index")
	}
	if testAdjMatrixDirected.AdjMatrix[1][2] != 0 {
		fmt.Printf("Data not found at index")
	}
}

func TestAdjacencyMatrix_AddEdgeUndirected() {
	testAdjMatrixUnDirected.Init()
	err := testAdjMatrixUnDirected.AddEdge(2, 1)
	if err != nil {
		fmt.Printf("Error adding edge")
	}
	if testAdjMatrixUnDirected.AdjMatrix[2][1] != 1 {
		fmt.Printf("Data not found at index")
	}
	if testAdjMatrixUnDirected.AdjMatrix[1][2] != 1 {
		fmt.Printf("Data not found at index")
	}
}

func TestAdjacencyMatrix_AddEdgeWeightDirectedGreaterThanVertex() {
	testAdjMatrixDirected.Init()
	err := testAdjMatrixDirected.AddEdgeWithWeight(10, 7, -3)
	if err == nil {
		fmt.Printf("Index was out of bounds it should have failed")
	}
}

func TestAdjacencyMatrix_AddEdgeWithWeightDirected() {
	testAdjMatrixDirected.Init()
	err := testAdjMatrixDirected.AddEdgeWithWeight(2, 1, -3)
	if err != nil {
		fmt.Printf("Error adding edge")
	}
	if testAdjMatrixDirected.AdjMatrix[2][1] != -3 {
		fmt.Printf("Data not found at index")
	}
	if testAdjMatrixDirected.AdjMatrix[1][2] != 0 {
		fmt.Printf("Data not found at index")
	}
}

func TestAdjacencyMatrix_AddEdgeWithWeightUnDirected() {
	testAdjMatrixUnDirected.Init()
	err := testAdjMatrixUnDirected.AddEdgeWithWeight(2, 1, -3)
	if err != nil {
		fmt.Printf("Error adding edge")
	}
	if testAdjMatrixUnDirected.AdjMatrix[2][1] != -3 {
		fmt.Printf("Data not found at index")
	}
	if testAdjMatrixUnDirected.AdjMatrix[1][2] != -3 {
		fmt.Printf("Data not found at index")
	}
}

func TestAdjacencyMatrix_RemoveEdgeDirected() {
	testAdjMatrixDirected.Init()
	testAdjMatrixDirected.AddEdge(2, 1)
	err := testAdjMatrixDirected.RemoveEdge(2, 1)
	if err != nil {
		fmt.Printf("Error removing edge")
	}
	if testAdjMatrixDirected.AdjMatrix[2][1] != 0 {
		fmt.Printf("Data not removed at index")
	}
}

func TestAdjacencyMatrix_RemoveEdgeUnDirected() {
	testAdjMatrixUnDirected.Init()
	testAdjMatrixUnDirected.AddEdge(2, 1)
	err := testAdjMatrixUnDirected.RemoveEdge(2, 1)
	if err != nil {
		fmt.Printf("Error removing edge")
	}
	if testAdjMatrixUnDirected.AdjMatrix[2][1] != 0 {
		fmt.Printf("Data not removed at index")
	}
	if testAdjMatrixUnDirected.AdjMatrix[1][2] != 0 {
		fmt.Printf("Data not removed at index")
	}
}

func TestAdjacencyMatrix_HasEdgeDirected() {
	testAdjMatrixDirected.Init()
	testAdjMatrixDirected.AddEdge(2, 1)
	if !testAdjMatrixDirected.HasEdge(2, 1) {
		fmt.Printf("No relationship, when there should be one")
	}
	if testAdjMatrixDirected.HasEdge(1, 2) {
		fmt.Printf("Relationship, when there shouldn't be one")
	}
}

func TestAdjacencyMatrix_HasEdgeUndirected() {
	testAdjMatrixUnDirected.Init()
	testAdjMatrixUnDirected.AddEdge(2, 1)
	if !testAdjMatrixUnDirected.HasEdge(2, 1) {
		fmt.Printf("No relationship, when there should be one")
	}
	if !testAdjMatrixUnDirected.HasEdge(1, 2) {
		fmt.Printf("No relationship, when there should be one")
	}
}

func TestAdjacencyMatrix_GetGraphType() {
	testAdjMatrixUnDirected.Init()
	if testAdjMatrixUnDirected.GraphType != testAdjMatrixUnDirected.GetGraphType() {
		fmt.Printf("GraphType not matching")
	}
}

func TestAdjacencyMatrix_GetAdjacentNodesForVertexDirectedOutOfBounds() {
	testAdjMatrixDirected.Init()
	testAdjMatrixDirected.AddEdge(2, 1)
	testAdjMatrixDirected.AddEdge(2, 0)
	testAdjMatrixDirected.AddEdge(1, 3)
	testAdjMatrixDirected.AddEdge(2, 2)
	testAdjMatrixDirected.AddEdge(2, 3)
	vertices := testAdjMatrixDirected.GetAdjacentNodesForVertex(10)
	if len(vertices) != 0 {
		fmt.Printf("Vertices should be 0. Data out of bounds")
	}
}

func TestAdjacencyMatrix_GetAdjacentNodesForVertexDirected() {
	testAdjMatrixDirected.Init()
	testAdjMatrixDirected.AddEdge(2, 1)
	testAdjMatrixDirected.AddEdge(2, 0)
	testAdjMatrixDirected.AddEdge(1, 1)
	testAdjMatrixDirected.AddEdge(2, 3)
	nodes := testAdjMatrixDirected.GetAdjacentNodesForVertex(2)
	if len(nodes) != 3 {
		fmt.Printf("Nodes size not matching. Size %d instead of %d. %v", len(nodes), 3, nodes)
	}
}

func TestAdjacencyMatrix_GetWeightOfEdgeDirected() {
	testAdjMatrixDirected.Init()
	testAdjMatrixDirected.AddEdgeWithWeight(2, 1, -3)
	weight, _ := testAdjMatrixDirected.GetWeightOfEdge(2, 1)
	if weight != -3 {
		fmt.Printf("Data not found at index")
	}
}

func TestAdjacencyMatrix_GetWeightOfEdgeUndirected() {
	testAdjMatrixUnDirected.Init()
	testAdjMatrixUnDirected.AddEdgeWithWeight(2, 1, -3)
	weight, _ := testAdjMatrixUnDirected.GetWeightOfEdge(2, 1)
	if weight != -3 {
		fmt.Printf("Data not found at index")
	}
}

func TestAdjacencyMatrix_GetNumberOfVertices() {
	testAdjMatrixUnDirected.Init()
	if testAdjMatrixUnDirected.GetNumberOfVertices() != testAdjMatrixUnDirected.Vertices {
		fmt.Printf("Vertex count doesn't match")
	}
}

func TestAdjacencyMatrix_GetNumberOfEdges() {
	testAdjMatrixDirected.Init()
	testAdjMatrixDirected.AddEdgeWithWeight(2, 1, -3)
	if testAdjMatrixDirected.GetNumberOfEdges() != 1 {
		fmt.Println("Edges count doesn't match", testAdjMatrixDirected.GetNumberOfEdges(), testAdjMatrixDirected.AdjMatrix)
	}
}

func TestAdjacencyMatrix_GetIndegreeForVertex() {
	testAdjMatrixDirected.Init()
	testAdjMatrixDirected.AddEdge(2, 1)
	testAdjMatrixDirected.AddEdge(2, 0)
	testAdjMatrixDirected.AddEdge(1, 1)
	testAdjMatrixDirected.AddEdge(2, 3)
	testAdjMatrixDirected.AddEdge(0, 3)
	if testAdjMatrixDirected.GetIndegreeForVertex(2) != 3 {
		fmt.Printf("Nodes size not matching.")
	}
}
