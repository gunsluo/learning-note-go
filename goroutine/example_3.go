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
	i := 0

L:
	for {
		select {
		case <-c:
			i++
			fmt.Println("i", i)
			if i == 2 {
				break L
			}
		}
	}

}

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}
