package main

import "fmt"

func main() {
	fmt.Println("gotest2")

	p := NewPerson("John", 45)
	p.Talk()

	d := NewDog("Fido", 5)
	d.Talk()
}
