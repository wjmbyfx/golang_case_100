package main

import (
	"fmt"
	"math"
)

/*
题目：一个整数，它加上100后是一个完全平方数，再加上168又是一个完全平方数，请问该数是多少？
*/

func main() {
	fmt.Println(solution())
}

func solution() (answer []int) {
	var res []int = make([]int, 0)
	for i := -100; i < 1000; i++ {
		if math.Sqrt(float64(i+100))-float64(int(math.Sqrt(float64(i+100)))) == 0 {
			if math.Sqrt(float64(i+168))-float64(int(math.Sqrt(float64(i+168)))) == 0 {
				res = append(res, i)
			}
		}
	}
	return res
}
