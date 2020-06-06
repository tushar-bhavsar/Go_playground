package main

import "fmt"

type Tushar int

var x Tushar
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%T\n", x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
