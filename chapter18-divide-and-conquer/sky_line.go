package main

import (
	"fmt"
)

// A Building is a rectangle on the skyline grounded at 0 elevations.
type Building struct {
	Left   int
	Right  int
	Height int
}

// type used to enable sorting a set of buildings
type Buildings []Building

func (b Buildings) Len() int {
	return len(b)
}

// Return true if the left edge of building at index i
// comes before the left edge of building at index j
func (b Buildings) Less(i, j int) bool {
	return b[i].Left < b[j].Left
}
func (b Buildings) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// A KeyPoint is a region on the skyline where the height changes.
type KeyPoint struct {
	X int
	Y int
}

type KeyPoints []*KeyPoint

func (c KeyPoints) Len() int {
	return len(c)
}
func (c KeyPoints) Less(i, j int) bool {
	// use greater to form a priority queue
	return c[i].Y > c[j].Y
}
func (c KeyPoints) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *KeyPoints) Push(x interface{}) {
	item := x.(*KeyPoint)
	*c = append(*c, item)
}

func (c *KeyPoints) Pop() interface{} {
	old := *c
	n := len(old)
	item := old[n-1]
	*c = old[0 : n-1]
	return item
}

func (c *KeyPoints) Peak() interface{} {
	old := *c
	return old[0]
}

func (c *KeyPoints) Max() int {
	if len(*c) == 0 {
		return 0
	}
	return c.Peak().(*KeyPoint).Y
}

// Function Signature for a function that can solve the skylines problem
type Solver func(Buildings) []KeyPoint

func (p KeyPoint) String() string {
	return fmt.Sprintf("KeyPoint(%d,%d)", p.X, p.Y)
}

func (p1 KeyPoint) Equals(p2 KeyPoint) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func skyline(buildings Buildings) (points []KeyPoint) {
	//first determine the width of the skyline
	width := 0
	for _, b := range buildings {
		if b.Right > width {
			width = b.Right
		}
	}

	// determine the maximum height at each point
	auxHeights := make([]int, width+1)
	for _, b := range buildings {
		for i := b.Left; i <= b.Right; i++ {
			if auxHeights[i] < b.Height {
				auxHeights[i] = b.Height
			}
		}
	}

	// read off the critical points from the skyline
	last_height := 0
	for i, h := range auxHeights {
		if h < last_height {
			points = append(points, KeyPoint{X: i - 1, Y: h})
		} else if h > last_height {
			points = append(points, KeyPoint{X: i, Y: h})
		}
		last_height = h
	}
	points = append(points, KeyPoint{X: width, Y: 0})
	return
}

func main() {
	buildings := []Building{Building{2, 9, 10}, Building{3, 7, 15}, Building{5, 12, 12}, Building{15, 20, 10}, Building{19, 24, 8}}
	fmt.Println(skyline(buildings))
}
