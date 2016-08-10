package main

import (
	"fmt"
	"time"
)

func main() {
	go ready("hello", 2)
	go ready("world", 1)

	fmt.Println("waiting!")
	time.Sleep(5 * time.Second)
}

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
}
