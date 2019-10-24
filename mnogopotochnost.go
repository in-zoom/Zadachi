package main

import (
	"fmt"
	"time"
)

//Напишите собственную функцию Sleep, используя time.After
func sleep() time.Time {

	select {
	case c2 := <-time.After(time.Second):
		return c2
	}
}
func main() {
	c1 := make(chan string)

	go func() {
		for {
			c1 <- "Привет"
			sleep()
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Сообщение 1", msg1)
				fmt.Println(sleep())
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
