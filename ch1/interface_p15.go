package main

import (
	"fmt"
)

type user struct {
	name string
	age  byte
}



func (u user) Print() {
	fmt.Printf("%+v\n", u)
}

type Printer interface {
	Print()
}

func main() {
	var u user
	u.name = "wq"
	u.age = 18
	var p Printer = u
	p.Print()
}

