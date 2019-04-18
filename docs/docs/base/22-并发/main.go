package main

import (
	"fmt"
	// "time"
)

func main() {
	// go say("666")
	// time.Sleep(2 * time.Second)
	// say("777")
	demo1()
}

func say(s string) {
	fmt.Println(s)
}

// channel demo
func demo1() {
	ch1 := make(chan int, 100)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	ch2 := <-ch1
	fmt.Printf("%v \n", ch2)
}
