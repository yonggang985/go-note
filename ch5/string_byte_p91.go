package main

import "fmt"

func main(){
	m := map[string]int{
		"abc":123,
	}


	key := []byte("abc")
	x,ok := m[string(key)]


	fmt.Println(x,ok)
}