package main

import "fmt"

func main() {
	data := [3]string{"a", "b", "c"}

	for i := range data {
		fmt.Println(i, data[i])
	}

	for k, v := range data {
		fmt.Println(k,v)
		fmt.Println(&k,&v)

	}

	for range data {

	}
}
