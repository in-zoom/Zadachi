package main

import "fmt"

//Напишите функцию, которая принимает число, делит его пополам и возвращает true в случае, если исходное число чётное, и false, если нечетное. Например, half(1) должна вернуть (0, false), в то время как half(2) вернет (1, true).

func chislo(xs int) int {
	//y := 2
	znach := xs / 2
	if znach%2 == 0 {
		return 1
	}
	return 0
}
func main() {
	xs := 10
	fmt.Println(chislo(xs))
}
