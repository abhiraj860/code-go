package main

import "fmt"

type Person struct {
	name string
	age int
}

func initPerson() *Person {
	p := Person{name:"Abhiraj", age: 56}
	fmt.Printf("%p\n", &p)
	return &p
}

func main() {
	fmt.Printf("%p\n", initPerson())
}