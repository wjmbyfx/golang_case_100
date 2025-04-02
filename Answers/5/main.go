package main

import (
	"errors"
	"fmt"
)

/*
题目：输入三个整数 x、y、z，请把这三个数由小到大输出。
*/

func main() {
	res, err := solution()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}

func solution() (res [3]int, err error) {
	var x, y, z int
	fmt.Scanf("%d %d %d", &x, &y, &z)
	if x == 0 || y == 0 || z == 0 {
		return [3]int{-1, -1, -1}, errors.New("Invalid input")
	}
	if x > y {
		x, y = y, x
	}
	if y > z {
		y, z = z, y
	}
	return [3]int{x, y, z}, nil
}
