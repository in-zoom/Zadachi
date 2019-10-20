package main

import "fmt"

/*Напишите программу, которая меняет местами
два числа (x := 1; y := 2; swap(&x, &y)
должно дать x=2 и y=1).*/
func swap(x *int, y *int) {
	*x = 2
	*y = 1

}
func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("x =", x, "|", "y =", y)
}
