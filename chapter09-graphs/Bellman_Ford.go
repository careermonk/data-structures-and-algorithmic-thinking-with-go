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

import "fmt"

import (
	"errors"
)

const inf = int(^uint(0) >> 1)

var ErrNegativeCycle = errors.New("bellman: graph contains negative cycle")
var ErrNoPath = errors.New("bellman: no path to vertex")

type Edges []*Edge

// Pair of vertices.
type Edge struct {
	From, To string
	Distance int
}

type Vertices map[string]*Vertex

// Vertex with predecessor information.
type Vertex struct {
	Distance int
	From     string // predecessor
}

// ShortestPath finds shortest path from source to destination after calling Search.
func (paths Vertices) ShortestPath(src, dest string) ([]*Vertex, error) {
	if len(paths) == 0 {
		// Did you forget calling Search?
		return nil, ErrNoPath
	}

	var all []*Vertex
	pred := dest

	for pred != src {
		if v := paths[pred]; v != nil {
			pred = v.From
			all = append(all, v)
		} else {
			return nil, ErrNoPath
		}
	}
	return all, nil
}

// AddEdge for search.
func (rt *Edges) AddEdge(from, to string, distance int) {
	*rt = append(*rt, &Edge{From: from, To: to, Distance: distance})
}

// Search for single source shortest path.
func (rt Edges) Search(start string) (Vertices, error) {
	// Resulting table constains vertex name to Vertex struct mapping.
	// Use v.From predecessor to trace the path back.
	distances := make(Vertices)
	for _, p := range rt {
		distances[p.From] = &Vertex{inf, " "}
		distances[p.To] = &Vertex{inf, " "}
	}
	distances[start].Distance = 0

	// As many iterations as there are nodes.
	for i := 0; i < len(distances); i++ {
		var changed bool

		// Iterate over pairs.
		for _, pair := range rt {
			n := distances[pair.From]
			if n.Distance != inf {
				if distances[pair.To].Distance > n.Distance+pair.Distance {
					distances[pair.To].Distance = n.Distance + pair.Distance
					distances[pair.To].From = pair.From
					changed = true
				}
			}
		}

		if !changed {
			// No more changes made. We done.
			break
		}
	}

	for _, pair := range rt {
		if distances[pair.From].Distance != inf &&
			distances[pair.To].Distance != inf &&
			distances[pair.From].Distance+pair.Distance < distances[pair.To].Distance {
			return nil, ErrNegativeCycle
		}
	}
	return distances, nil
}

func main() {
	TestSearch()
	TestShortestPath()
}

func TestSearch() {
	rt := Edges{
		{"A", "E", -3},
		{"B", "C", 7},
		{"C", "F", -5},
		{"E", "B", 1},
		{"E", "G", 6},
		{"E", "H", -3},
		{"F", "B", 8},
		{"G", "D", 2},
		{"H", "G", 1},
		{"S", "A", 2},
	}

	distances, err := rt.Search("S")
	if err != nil {
		fmt.Println(err)
	}

	distance := map[string]int{
		"S": 0, "A": 2, "B": 0,
		"C": 7, "D": -1, "E": -1,
		"F": 2, "G": -3, "H": -4,
	}
	// list of predecessors
	pred := map[string]string{
		"S": " ", "A": "S", "B": "E",
		"C": "B", "D": "G", "E": "A",
		"F": "C", "G": "H", "H": "E",
	}

	for b, n := range distances {
		if distance[b] != n.Distance {
			fmt.Printf("route %q distance got %d want %d", b, n.Distance, distance[b])
		}
		if pred[b] != n.From {
			fmt.Printf("route %q predecessor got %q want %q", b, n.From, pred[b])
		}
	}
}

type testcase struct {
	src, dest string
	path      string
}

var tests = []*testcase{
	&testcase{"Hyderabad", "Chennai", "Hyderabad 106,Bangalore 95,"},
	&testcase{"Bangalore", "Mumbai", "Bangalore 95,Chennai 65,Canada 73,Pune 62,Amaravathi 92,"},
	&testcase{"Delhi", "Chennai", "Delhi 121,"},
}

func TestShortestPath() {
	for _, test := range tests {
		rt := Edges{
			{"Hyderabad", "Vijayawada", 117},
			{"Vijayawada", "Vizag", 65},
			{"Vizag", "Pune", 70},
			{"Pune", "Amaravathi", 62},
			{"Amaravathi", "Mumbai", 92},
			{"Hyderabad", "Bangalore", 106},
			{"Hyderabad", "Bangalore", 113},
			{"Bangalore", "Chennai", 95},
			{"Vijayawada", "Guntur", 24},
			{"Guntur", "Chennai", 88},
			{"Vizag", "Chennai", 94},
			{"Chennai", "Canada", 65},
			{"Canada", "Pune", 73},
			{"Chennai", "Delhi", 121},
			{"Canada", "Delhi", 103},
		}
		for _, r := range rt {
			// Add reversed edges.
			rt = append(rt, &Edge{From: r.To, To: r.From, Distance: r.Distance})
		}

		paths, err := rt.Search(test.src)
		if err != nil {
			fmt.Println(err)
		}

		all, err := paths.ShortestPath(test.src, test.dest)
		if err != nil {
			fmt.Println(err)
		}

		var prev int
		var s string
		for i := len(all) - 1; i >= 0; i-- {
			s += fmt.Sprintf("%s %d,", all[i].From, all[i].Distance-prev)
			prev = all[i].Distance
		}

		if test.path != s {
			fmt.Printf("shortest path does not match, got %q want %q", s, test.path)
		}
	}
}
