package main

import "fmt"

/*Последовательность чисел Фибоначчи определяется
как fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
Напишите рекурсивную функцию, находящую fib(n).*/
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)

}
func main() {
	//n := 12
	fmt.Println(fib(12))

}
