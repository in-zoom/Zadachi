package main

import "fmt"

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
