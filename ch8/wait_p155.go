package main

import (
	"time"
	"fmt"
)

func main(){
	exit := make(chan struct{})

	go func(){
		time.Sleep(time.Second)
		fmt.Println("goroutine done.")

		close(exit)
	}()

	fmt.Println("main ...")
	<-exit
	fmt.Println("main exit.")
}