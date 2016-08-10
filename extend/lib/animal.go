package lib

import "fmt"

type Animal struct {
	age int
}

func (this *Animal) SetAge(age int) {
	this.age = age
}

func (this *Animal) Age() int {
	return this.age
}

func (this *Animal) say() {
	fmt.Printf("say: I'am %d wa wa wa !!!\n", this.age)
}

func (this *Animal) Say() {
	fmt.Printf("Say: I'am %d wa wa wa !!!\n", this.age)
}
