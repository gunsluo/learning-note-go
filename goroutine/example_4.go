package main

import (
	"fmt"
	"runtime"
	"time"
)

var c chan int

func main() {

	runtime.GOMAXPROCS(2)

	c = make(chan int)

	go send()
	go loop()

	var i int

	for {
		select {
		case <-c:
			i++
			fmt.Println("i", i)
		}
	}
}

func send() {
	fmt.Println("send!")
	c <- 1
}

func loop() {

	for {
		time.Sleep(time.Second * 30)
	}
}
