package main

import (
	"fmt"
	"time"
)

//Напишите собственную функцию Sleep, используя time.After
func sleep(s int) {

	select {
	case <-time.After(time.Second * time.Duration(s)):
		return
	}
}

func main() {
	c1 := make(chan string)

	go func() {
		for {
			c1 <- "Привет"
			sleep(3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Сообщение 1", msg1)

			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
