package main

import (
	"fmt"

	"github.com/gunsluo/learning-note-go/extend/lib"
)

type Person struct {
	lib.Animal
}

func (this *Person) Say() {
	fmt.Printf("Say: I'am %d, hello world !!!\n", this.Age())
}

func main() {
	animal := new(lib.Animal)
	animal.SetAge(10)

	animal.Say()

	person := new(Person)
	person.SetAge(20)
	//person.say()
	person.Say()
}
