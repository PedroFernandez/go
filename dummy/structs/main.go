package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	fmt.Println(Person{"Pedro", 42})
	newPerson := Person{"Marta", 4}
	fmt.Println(newPerson)
	newPerson.age = 5
	fmt.Println(newPerson)
}