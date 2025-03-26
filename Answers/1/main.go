package main

import "fmt"

/*

有 1、2、3、4 四个数字，能组成多少个互不相同且无重复数字的三位数？
都是多少？

*/

func main() {
	fmt.Println(solution())
}

func solution() (int, []int) {
	var result []int = make([]int, 0)
	var count int = 0
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			for k := 1; k <= 4; k++ {
				if i != j && j != k && i != k {
					result = append(result, i*100+j*10+k)
					count++
				}
			}
		}
	}
	return count, result
}
