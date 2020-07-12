/*
# Copyright (c) July 12, 2020 CareerMonk Publications and others.
# E-Mail           		: info@careermonk.com
# Creation Date    		: 2020-07-12 06:15:46
# Last modification		: 2020-07-12
#               by		: Narasimha Karumanchi
# Book Title			: Data Structures And Algorithmic Thinking With Go
# Warranty         		: This software is provided "as is" without any
# 				   warranty; without even the implied warranty of
# 				    merchantability or fitness for a particular purpose. */

package main
import "fmt"

func TOHUtil(n int, from, to, temp string) {
	/* If only 1 disk, make the move and return */
	if n == 1 {
		fmt.Println("Move disk ", n, " from peg ", from, " to peg ", to)
		return
	}
	/* Move top n-1 disks from A to B, using C as auxiliary */
	TOHUtil(n-1, from, temp, to)
	/* Move remaining disks from A to C */
	fmt.Println("Move disk ", n, " from peg ", from, " to peg ", to)
	/* Move n-1 disks from B to C using A as auxiliary */
	TOHUtil(n-1, temp, to, from)
}
func TowersOfHanoi(n int) {
	TOHUtil(n, "A", "C", "B")
}
func main() {
	TowersOfHanoi(3)
}
