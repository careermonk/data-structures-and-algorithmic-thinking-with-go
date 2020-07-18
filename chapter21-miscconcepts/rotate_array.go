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

func leftRotate(arr []int, d int, n int) {
	if d == 0 {
		return
	}
	reverseArray(arr, 0, d-1)
	reverseArray(arr, d, n-1)
	reverseArray(arr, 0, n-1)
}

func printArray(arr []int, size int) {
	var i int
	//UTILITY FUNCTIONS
	// function to print an array
	for i = 0; i < size; i++ {
		fmt.Printf("%d ", arr[i])
	}
}

func reverseArray(arr []int, start int, end int) {
	var temp int
	for start < end {
		//Function to reverse arr[] from index start to end
		temp = arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}

// Driver program to test above functions
func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7}
	var n = len(arr)
	var d = 2
	// in case the rotating factor is greater than array length
	d = d % n
	leftRotate(arr, d, n)
	printArray(arr, n)
	return
}
