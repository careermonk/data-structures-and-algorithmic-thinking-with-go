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

func swap(num int) int {
	even := 0xAAAAAAAA
	odd := 0x55555555
	return (num << 1 & even) | (num >> 1 & odd)
}

func main() {
	fmt.Println(swap(5))
	fmt.Println(swap(10))
	fmt.Println(swap(2))
	fmt.Println(swap(4))
}
