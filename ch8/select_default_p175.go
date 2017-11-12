package main

import "fmt"

func main() {
	done := make(chan struct{})
	data := []chan int{
		make(chan int, 3),
	}

	go func() {
		defer close(done)

		for i := 0; i < 10; i++ {
			select {
			case data[len(data)-1] <- i:
				fmt.Println("len data:",len(data))
			default:
				data = append(data, make(chan int, 3))
			}
		}
	}()

	<-done

	for i := 0; i < len(data); i++ {
		fmt.Println("len",len(data))
		c := data[i]
		close(c)

		for x := range c {
			fmt.Println(x)
		}
		//close(c)
	}
}
