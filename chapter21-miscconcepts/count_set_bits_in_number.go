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

func countSetBits1(n int) int {
	var count int
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func countSetBits2(n int) int {
	var count int
	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n = n / 2
	}
	return count
}

func countSetBits3(n int) int {
	var count int
	for n > 0 {
		count++
		n &= n - 1
	}
	return count
}


var Table = []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}	
func countSetBits4(n int) int {
	count := 0
	for ; n > 0; n >>= 4 {
		count = count + Table[n&0xF]
		return count
	}
	return count
}

// Program to test function countSetBits
func main() {
	var i = 9
	fmt.Println(countSetBits1(i))
	fmt.Println(countSetBits2(i))
	fmt.Println(countSetBits3(i))
	fmt.Println(countSetBits4(i))
}
