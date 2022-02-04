package main

import (
	"fmt"
)

func main() {
	// (type, capacity)
	c := make(chan string, 2)
	
	c <- "hello1"
	c <- "hello2"

	msg := <- c
	fmt.Println(msg)

	msg = <- c
	fmt.Println(msg)
}
