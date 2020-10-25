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
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findConnects(matrix [][]int, M, N, r, c int) int {
	answer := 0
	if r < 0 || c < 0 || r >= M || c >= N {
		answer = 0

	} else if matrix[r][c] == 1 {
		matrix[r][c] = 0
		answer = 1 +
			findConnects(matrix, M, N, r-1, c) +
			findConnects(matrix, M, N, r+1, c) +
			findConnects(matrix, M, N, r, c-1) +
			findConnects(matrix, M, N, r, c+1) +
			findConnects(matrix, M, N, r-1, c-1) +
			findConnects(matrix, M, N, r-1, c+1) +
			findConnects(matrix, M, N, r+1, c-1) +
			findConnects(matrix, M, N, r+1, c+1)
	}
	return answer
}

func findMaxConnects(matrix [][]int, M, N int) int {
	maxConnects := 0
	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			maxConnects = max(maxConnects, findConnects(matrix, M, N, r, c))
		}
	}
	return maxConnects
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	matrix := make([][]int, M)
	for i := 0; i < M; i++ {
		matrix[i] = make([]int, N)
		for j := 0; j < N; j++ {
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			matrix[i][j] = v
		}
	}

	fmt.Println(findMaxConnects(matrix, M, N))
}
