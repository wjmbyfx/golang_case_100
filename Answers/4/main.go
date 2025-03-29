package main

import (
	"errors"
	"fmt"
)

/*
输入某年某月某日，判断这一天是这一年的第几天？
*/

var leap_year_map [12]int = [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var common_year_map [12]int = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func main() {
	fmt.Println(solution())
}

func solution() (res int, err error) {
	var year, month, day int
	fmt.Scanf("%d-%d-%d", &year, &month, &day)
	if year < 0 || month <= 0 || month > 12 || day <= 0 || day > 31 {
		return -1, errors.New("In valid input")
	}
	var total = 0
	if is_leap_year(year) {
		for i := 0; i < month-1; i++ {
			total += leap_year_map[i]
		}

	} else {
		for i := 0; i < month-1; i++ {
			total += common_year_map[i]
		}

	}
	total += day
	return total, nil
}

func is_leap_year(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}
