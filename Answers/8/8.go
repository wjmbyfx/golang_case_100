package main

import "fmt"

/*
题目：输出 9*9 口诀。
*/

func main() {
	solution()
}

func solution() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
