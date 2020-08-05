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
	"fmt"
)

type Stmt struct {
	Name  string
	After []string
}

func main() {
	stmts := []Stmt{
		{Name: "app", After: []string{"app_user"}},
		{Name: "billingplan", After: []string{}},
		{Name: "campaign", After: []string{"app_user"}},
		{Name: "campaign_app", After: []string{"campaign", "app"}},
		{Name: "campaign_ip", After: []string{"campaign", "ip"}},
		{Name: "campaign_operator", After: []string{"campaign", "operator"}},
		{Name: "campaign_sponsor", After: []string{"campaign", "sponsor"}},
		{Name: "campaign_subscriberfilter", After: []string{"campaign", "subscriber_filters"}},
		{Name: "campaign_url", After: []string{"campaign", "url"}},
		{Name: "contentpartner", After: []string{"app_user"}},
		{Name: "filter_criteria", After: []string{"campaign", "subscriber_filters"}},
		{Name: "ip", After: []string{"app_user"}},
		{Name: "mobile_registered", After: []string{"campaign", "app"}},
		{Name: "operator", After: []string{}},
		{Name: "passwords", After: []string{"app_user"}},
		{Name: "publish_package", After: []string{}},
		{Name: "role", After: []string{}},
		{Name: "passwords", After: []string{"app_user"}},
		{Name: "sponsor", After: []string{"app_user"}},
		{Name: "subscriber_dbs", After: []string{}},
		{Name: "subscriber_filters", After: []string{"subscriber_dbs"}},
		{Name: "timezone", After: []string{}},
		{Name: "url", After: []string{"app_user"}},
		{Name: "app_user", After: []string{}},
		{Name: "user_role", After: []string{"app_user", "role"}},
	}

	g := make(graph)
	for _, s := range stmts {
		g[s.Name] = after(s.After)
	}

	t := newTarjan(g)
	for _, s := range t.sccs {
		fmt.Println(s)
	}
}

// graph is an edge list representation of a graph.
type graph map[string]set

// set is an integer set.
type set map[string]struct{}

func after(i []string) set {
	if len(i) == 0 {
		return nil
	}
	s := make(set)
	for _, v := range i {
		s[v] = struct{}{}
	}
	return s
}

// tarjan implements Tarjan's strongly connected component finding
// algorithm. The implementation is from the pseudocode at
//
// http://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm
//
type tarjan struct {
	g graph

	index      int
	indexTable map[string]int
	lowLink    map[string]int
	onStack    map[string]bool

	stack []string

	sccs [][]string
}

// newTarjan returns a tarjan with the sccs field filled with the
// strongly connected components of the directed graph g.
func newTarjan(g graph) *tarjan {
	t := tarjan{
		g: g,

		indexTable: make(map[string]int, len(g)),
		lowLink:    make(map[string]int, len(g)),
		onStack:    make(map[string]bool, len(g)),
	}
	for v := range t.g {
		if t.indexTable[v] == 0 {
			t.strongconnect(v)
		}
	}
	return &t
}

// strongconnect is the strongconnect function described in the
// wikipedia article.
func (t *tarjan) strongconnect(v string) {
	// Set the depth index for v to the smallest unused index.
	t.index++
	t.indexTable[v] = t.index
	t.lowLink[v] = t.index
	t.stack = append(t.stack, v)
	t.onStack[v] = true

	// Consider successors of v.
	for w := range t.g[v] {
		if t.indexTable[w] == 0 {
			// Successor w has not yet been visited; recur on it.
			t.strongconnect(w)
			t.lowLink[v] = min(t.lowLink[v], t.lowLink[w])
		} else if t.onStack[w] {
			// Successor w is in stack s and hence in the current SCC.
			t.lowLink[v] = min(t.lowLink[v], t.indexTable[w])
		}
	}

	// If v is a root node, pop the stack and generate an SCC.
	if t.lowLink[v] == t.indexTable[v] {
		// Start a new strongly connected component.
		var (
			scc []string
			w   string
		)
		for {
			w, t.stack = t.stack[len(t.stack)-1], t.stack[:len(t.stack)-1]
			t.onStack[w] = false
			// Add w to current strongly connected component.
			scc = append(scc, w)
			if w == v {
				break
			}
		}
		// Output the current strongly connected component.
		t.sccs = append(t.sccs, scc)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
