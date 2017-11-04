package main

import (
	"fmt"
	"runtime"
)

func producer(data chan int){
	for i := 0;i<10;i++{
		data <-i
	}
	fmt.Println("producer")
	close(data)
}

func comsumer(data chan int, done chan bool){
	for i := range data {
		fmt.Println("recv:",i)
	}
	fmt.Println("comsumer")
	done <- true
}

func main(){
	runtime.GOMAXPROCS(4)
	done := make(chan bool)
	data := make(chan int)

	go comsumer(data,done)
	go producer(data)
	<- done
}