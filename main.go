package main

import (
	"Zadachi/math"
	"Zadachi/max"
	"Zadachi/min"
	"fmt"
)

/*Мы скопировали функцию Average из главы 7 в наш новый пакет.
Создайте Min и Max функции для нахождения наименьших и наибольших значений в срезах дробных чисел типа float64.*/
func main() {
	xs := []float64{0.9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	avg := math.Average(xs)
	avmin := min.Min(xs)
	avmax := max.Max(xs)
	fmt.Println(avg, "|", "min=", avmin, "|", "max=", avmax)
}
