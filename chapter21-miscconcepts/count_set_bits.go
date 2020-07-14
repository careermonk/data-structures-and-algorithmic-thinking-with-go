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

func countSetBits(n int) int {
	count := 0
	i := 0
	var j int
	for i = 1; i <= n; i++ {
		j = i
		for j > 0 {
			j = j & (j - 1)
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countSetBits(5))
	fmt.Println(countSetBits(10))
	fmt.Println(countSetBits(2))
	fmt.Println(countSetBits(4))
}
