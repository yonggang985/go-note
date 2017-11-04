package main

import "fmt"

func main() {
	data := [3]int{10, 20, 30}

	for k, v := range data {
		if k == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("v: %d,data: %d\n", v, data[k])
	}

	for k, v := range data[:] {
		if k == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("v: %d,data: %d\n", v, data[k])
	}


