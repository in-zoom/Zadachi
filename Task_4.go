package main

import "fmt"

/*Используя в качестве примера функцию makeEvenGenerator
напишите makeOddGenerator, генерирующую нечётные числа.*/
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i++
		if i%2 == 0 {
			i++
		}
		return ret
	}
}
func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
