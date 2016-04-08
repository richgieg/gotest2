package main

import "fmt"

type Person struct {
	Animal
}

func NewPerson(name string, age int) *Person {
	p := new(Person)
	p.Name = name
	p.Age = age
	return p
}

func (p *Person) Talk() {
	fmt.Printf("Hi! My name is %s\n", p.Name)
}
