package main

import "fmt"

/*
题目：用*号输出字母C的图案。

*/

func main() {
	solution()
}

func solution() {
	var output_matrix []([]rune) = make([][]rune, 4)
	for i := 0; i < 4; i++ {
		output_matrix[i] = make([]rune, 4)
	}
	output_matrix[0] = []rune{'*', '*', '*', '*'}
	output_matrix[1] = []rune{'*', ' ', ' ', ' '}
	output_matrix[2] = []rune{'*', ' ', ' ', ' '}
	output_matrix[3] = []rune{'*', '*', '*', '*'}
	for i := range output_matrix {
		fmt.Println((output_matrix[i]))
	}
}
