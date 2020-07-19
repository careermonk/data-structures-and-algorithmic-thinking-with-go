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

import (
	"fmt"
	"sort"
)

type item struct {
	item   string
	weight float64
	price  float64
}

type items []item

var all = items{
	{"charger", 3.8, 36},
	{"rice", 5.4, 43},
	{"bread", 3.6, 90},
	{"grapes", 2.4, 45},
	{"apples", 4.0, 30},
	{"shirts", 2.5, 56},
	{"trousers", 3.7, 67},
	{"laptop", 3.0, 95},
	{"sausage", 5.9, 98},
}

// satisfy sort interface
func (z items) Len() int      { return len(z) }
func (z items) Swap(i, j int) { z[i], z[j] = z[j], z[i] }
func (z items) Less(i, j int) bool {
	return z[i].price/z[i].weight > z[j].price/z[j].weight
}

func main() {
	left := 15.
	sort.Sort(all)
	for _, i := range all {
		if i.weight <= left {
			fmt.Println("take all the", i.item)
			if i.weight == left {
				return
			}
			left -= i.weight
		} else {
			fmt.Printf("take %.1fkg %s\n", left, i.item)
			return
		}
	}
}
