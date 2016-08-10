package main

import "fmt"

type Set struct {
}

func (this *Set) Add(node interface{}) bool {
	fmt.Println("set add")
	return true
}

func (this *Set) Del(node interface{}) bool {
	fmt.Println("set del")
	return true
}

func (this *Set) List() {
	fmt.Println("set list")
}

func main() {
	var c Collection = new(Set)

	c.Add(1)
	c.Del(2)
	c.List()
}
