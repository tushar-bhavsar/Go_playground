package main

import "fmt"

type Tushar int

var x Tushar

func main() {
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
