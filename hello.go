package main

import (
	"fmt"
	"time"
)

func say(s int, c chan int) {
	fmt.Println("HELLO", s)
	time.Sleep(1 * time.Second)
}

func main() {

	for i := 1; i < 6; i++ {
		channel := make(chan int, i)
		go say(i, channel)
	}
	time.Sleep(1 * time.Second)
}
