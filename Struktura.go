package main

import (
	"fmt"
	"math"
)

/*Добавьте новый метод perimeter в интерфейс Shape,
который будет вычислять периметр фигуры.
Имплементируйте этот метод для Circle и Rectangle*/

type Shape interface {
	area() float64
	perimeter() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type Circle struct {
	x, y, r float64
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (r Rectangle) perimeter() float64 {
	h := r.x1 + r.x2 + r.y1 + r.y2
	return h
}

func main() {
	var a Shape = &Circle{0, 0, 5}
	var b Shape = &Rectangle{0, 0, 10, 10}
	fmt.Println(a.area())
	fmt.Println(b.area())
	fmt.Println("Периментр круга", a.perimeter())
	fmt.Println("Периметр прямрукольника", b.perimeter())

}
