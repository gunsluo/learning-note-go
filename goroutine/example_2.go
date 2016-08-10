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

	go ready("hello", 2)
	go ready("world", 1)

	fmt.Println("waiting!")
	<-c
	<-c
	//time.Sleep(5 * time.Second)
}

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}
