package main

import (
	"fmt"
	"time"
)

func say(s string, c chan int) {
	fmt.Println(s, c)
	time.Sleep(1 * time.Second)
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)
	chan4 := make(chan int)
	go say("hello 1", chan1)
	go say("hello 2", chan2)
	go say("hello 3", chan3)
	go say("hello 4", chan4)
	time.Sleep(1 * time.Second)
}
