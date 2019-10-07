package main

import "fmt"

func main() {
x := []int{
48, 96, 86, 68,
57, 1, 63, 70,
37, 34, 83, 27,
19, 97, 9, 17	,
}
y := []int{x[0]} 
for i, value := range x {
if value < y[0] {
y[0] = x[i] 
}
}
fmt.Println(y[0])
}
