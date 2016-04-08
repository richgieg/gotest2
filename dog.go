package main

import "fmt"

type Dog struct {
	Name string
	Age  int
}

func NewDog(name string, age int) *Dog {
	d := new(Dog)
	d.Name = name
	d.Age = age
	return d
}

func (d *Dog) Talk() {
	fmt.Printf("Woof! My name is %s\n", d.Name)
}
