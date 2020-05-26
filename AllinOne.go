package main

import (
	"fmt"
)

var a int
var b string

type aws int

var c aws
var d bool

const (
	e = 300
	f = iota
	g
	h
)

func main() {

	a = 100
	b = "This is test going on"
	c = 200
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

}
