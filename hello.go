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
	chan1 := make(chan int)
	for i := 1; i < 6; i++ {
		go say(i, chan1)
	}
	time.Sleep(1 * time.Second)
}
