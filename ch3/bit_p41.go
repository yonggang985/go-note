package main

import "fmt"

const (
	read byte = 1 << iota
	write
	exec
	freeze
)

func main(){
	a := read | write | freeze
	b := read | freeze | exec
	c := a &^ b

	fmt.Printf("%08b, %08b,%08b,%08b \n",read,write,exec,freeze)
	fmt.Printf("%08b &^ %08b = %08b\n",a,b,c)

}