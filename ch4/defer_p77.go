package main

import "fmt"

func test() (z int){
	defer func() {
		fmt.Println("defer:",z)
		z += 100
	}()
	return 100
}
func main()  {
	fmt.Println("test:",test())
}
