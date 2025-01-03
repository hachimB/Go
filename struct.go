package main

import "fmt"

type Sport struct {
	name string
	post string
}

type Person struct {
	name     string
	age      int
	favSport Sport
}

// Method
func (p Person) getNameFunction() string {
	return p.name
}

func structure() {
	p1 := Person{}
	p2 := Person{age: 23, name: "idk"}
	p1.age = 21
	p1.name = "Hachim"
	fmt.Println(p1, p2)

	fmt.Println(p2.getNameFunction())

	p3 := Person{}
	p3.name = "someone"
	p3.age = 23
	p3.favSport.post = "Striker"
	p3.favSport.name = "Soccer"
	fmt.Println(p3)
}
