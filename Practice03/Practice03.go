package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)

	fmt.Println(s)
}
