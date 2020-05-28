package main

import "fmt"

const (
	a        = 50
	b string = "This is Constant"
	c        = 79.40
)

func main() {

	fmt.Println(a, b, c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
