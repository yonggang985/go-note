package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 10
	c <- 20
	close(c)
	for i := 0; i < cap(c)+1; i++ {
		x, ok := <-c
		fmt.Println(i, ":", ok, x)
	}
}
