package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (this Person) GetName() string {
	return this.name
}

func (this *Person) Name() string {
	return this.name
}

func (this Person) GetAge() int {
	return this.age
}

func (this *Person) Age() int {
	return this.age
}

func main() {

	var i int = 10
	var pi *int = &i
	*pi = 20
	fmt.Printf("i=%d pi=%p *pi=%d\n", i, pi, *pi)

	var person1 = new(Person)
	person1.name = "luoji"
	person1.age = 20

	fmt.Printf("name=%s age=%d type=%T value=%p\n", person1.name, person1.age, person1, person1)
	fmt.Printf("name=%s age=%d type=%T value=%p\n", person1.Name(), person1.Age(), person1, person1)
	fmt.Printf("name=%s age=%d type=%T value=%p\n", person1.GetName(), person1.GetAge(), person1, person1)

	var person2 = Person{
		name: "jerrylou",
		age:  18,
	}

	fmt.Printf("name=%s age=%d type=%T value=%v\n", person2.name, person2.age, person2, person2)
	fmt.Printf("name=%s age=%d type=%T value=%v\n", person2.Name(), person2.Age(), person2, person2)
	fmt.Printf("name=%s age=%d type=%T value=%v\n", person2.GetName(), person2.GetAge(), person2, person2)
}
