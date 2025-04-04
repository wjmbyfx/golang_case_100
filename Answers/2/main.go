package main

import "fmt"

/*

企业发放的奖金根据利润提成。

利润(I)低于或等于10万元时，奖金可提10%；
利润高于10万元，低于20万元时，低于10万元的部分按10%提成，高于10万元的部分，可提成7.5%；
20万到40万之间时，高于20万元的部分，可提成5%；
40万到60万之间时高于40万元的部分，可提成3%；
60万到100万之间时，高于60万元的部分，可提成1.5%；
高于100万元时，超过100万元的部分按1%提成。
从键盘输入当月利润I，求应发放奖金总数？

*/

func main() {
	solution()
}

func solution() {
	var I float64
	var bonus float64
	fmt.Scanf("%f", &I)
	if I <= 10 {
		bonus = I * 0.1
	} else if I > 10 && I <= 20 {
		bonus = 1 + (I-10)*0.075
	} else if I > 40 && I <= 40 {
		bonus = 1 + (20-10)*0.075 + (I-20)*0.05
	} else if I > 60 && I <= 100 {
		bonus = 1 + (20-10)*0.075 + (60-20)*0.05 + (I-60)*0.015
	} else if I > 100 {
		bonus = 1 + (20-10)*0.075 + (60-20)*0.05 + (100-60)*0.015 + (I-100)*0.01
	}
	fmt.Println(bonus)
}
