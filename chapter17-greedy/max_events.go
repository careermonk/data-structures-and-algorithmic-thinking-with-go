// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

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
	"sort"
)

func maxEvents(start, finish []int) int {
	sort.Slice(start, func(i, j int) bool {
		if finish[i] == finish[j] {
			return start[i] < start[j]
		}

		return finish[i] < finish[j]
	})

	m := make(map[int]bool)

	for j := 0; j < len(start); j++ {
		for i := start[j]; i <= finish[j]; i++ {
			if _, ok := m[i]; !ok {
				m[i] = true
				break
			}
		}
	}

	return len(m)
}

// Driver program to test above functions
func main() {
	fmt.Println(maxEvents([]int{1, 2, 3}, []int{2, 3, 4}))
}
