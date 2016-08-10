package main

import "fmt"

type Animal struct {
	age int
}

func (this *Animal) say() {
	fmt.Printf("say: I'am %d wa wa wa !!!\n", this.age)
}

func (this *Animal) Say() {
	fmt.Printf("Say: I'am %d wa wa wa !!!\n", this.age)
}

type Person struct {
	Animal
}

func (this *Person) Say() {
	fmt.Printf("Say: I'am %d, hello world !!!\n", this.age)
}

func main() {
	animal := new(Animal)
	animal.age = 10

	animal.say()
	animal.Say()

	person := new(Person)
	person.age = 20
	person.say()
	person.Say()
}
