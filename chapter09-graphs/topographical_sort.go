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

// topographicalSort Input: G: a directed acyclic graph with vertices number 1..n
// Output: a linear order of the vertices such that u appears before v
// in the linear order if (u,v) is an edge in the graph.
func topographicalSort(G map[int][]int) []int {
	topoOrder := []int{}

	// 1. Let inDegree[1..n] be a new array, and create an empty linear array of verticies
	inDegree := map[int]int{}

	// 2. Set all values in inDegree to 0
	for u := range G {
		inDegree[u] = 0
	}

	// 3. For each vertex u
	for _, adjacent := range G {
		// A. For each vertex *v* adjacent to *u*:
		for _, v := range adjacent {
			//  i. increment inDegree[v]
			inDegree[v]++
		}
	}

	// 4. Make a list next consisting of all vertices u such that in-degree[u] = 0
	next := []int{}
	for u, v := range inDegree {
		if v != 0 {
			continue
		}
		next = append(next, u)
	}

	// 5. While next is not empty...
	for len(next) > 0 {
		// A. delete a vertex from next and call it vertex u
		u := next[0]
		next = next[1:]

		// B. Add u to the end of the linear order
		topoOrder = append(topoOrder, u)

		// C. For each vertex v adjacent to u
		for _, v := range G[u] {
			// i. Decrement inDegree[v]
			inDegree[v]--

			// ii. if inDegree[v] = 0, then insert v into next list
			if inDegree[v] == 0 {
				next = append(next, v)
			}
		}
	}

	// 6. Return the linear order
	return topoOrder
}
func main() {
	// Directed Acyclic Graph
	vertices := map[int][]int{
		1:  []int{4},
		2:  []int{3},
		3:  []int{4, 5},
		4:  []int{6},
		5:  []int{6},
		6:  []int{7, 11},
		7:  []int{8},
		8:  []int{14},
		9:  []int{10},
		10: []int{11},
		11: []int{12},
		13: []int{13},
		14: []int{},
	}

	// As yet unimplemented topographicalSort
	fmt.Println(topographicalSort(vertices))
}
