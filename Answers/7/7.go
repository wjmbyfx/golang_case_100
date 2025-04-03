package main

import "fmt"

/*
题目：输出特殊图案，请在c环境中运行，看一看，Very Beautiful!
*/

func main() {
	solution()
}

func solution() {
	var a, b rune = 176, 219
	fmt.Printf("%c%c%c%c%c\n", b, a, a, a, b)
	fmt.Printf("%c%c%c%c%c\n", a, b, a, b, a)
	fmt.Printf("%c%c%c%c%c\n", a, a, b, a, a)
	fmt.Printf("%c%c%c%c%c\n", a, b, a, b, a)
	fmt.Printf("%c%c%c%c%c\n", b, a, a, a, b)
}
