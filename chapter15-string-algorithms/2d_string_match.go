// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
//

package main

import (
	"fmt"
)

func findMatch(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	pos := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if helper(visited, board, i, j, word, pos) {
				return true
			}
		}
	}
	return false
}
func helper(visited [][]bool, board [][]byte, r, c int, word string, pos int) bool {
	if pos == len(word) {
		return true
	}
	if r < 0 || r == len(board) || c < 0 || c == len(board[0]) {
		return false
	}
	if board[r][c] != word[pos] || visited[r][c] {
		return false
	}
	//finding subpattern in 8 neighbors
	visited[r][c] = true
	if helper(visited, board, r-1, c, word, pos+1) {
		return true
	}
	if helper(visited, board, r+1, c, word, pos+1) {
		return true
	}
	if helper(visited, board, r, c-1, word, pos+1) {
		return true
	}
	if helper(visited, board, r, c+1, word, pos+1) {
		return true
	}
	if helper(visited, board, r+1, c+1, word, pos+1) {
		return true
	}
	if helper(visited, board, r+1, c-1, word, pos+1) {
		return true
	}
	if helper(visited, board, r-1, c+1, word, pos+1) {
		return true
	}
	if helper(visited, board, r-1, c-1, word, pos+1) {
		return true
	}
	visited[r][c] = false
	return false
}
func main() {
	board := [][]byte{
		{'A', 'C', 'P', 'R', 'C'},
		{'X', 'S', 'O', 'P', 'C'},
		{'V', 'O', 'V', 'N', 'I'},
		{'W', 'G', 'F', 'M', 'N'},
		{'Q', 'A', 'T', 'I', 'T'},
	}
	fmt.Println(findMatch(board, "MICROSOFT"))
}
