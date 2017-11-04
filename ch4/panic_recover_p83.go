package main

import (
	"log"
	"fmt"
)

func test() {
	defer fmt.Println("test.1")
	defer fmt.Println("test.2")

	panic("i am dead")
}

func main(){
	defer func() {
		log.Println(recover())
	}()

	test()
}