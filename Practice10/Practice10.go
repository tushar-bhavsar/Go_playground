package main

import "fmt"

func main() {
	for i := 10; i <= 15; i++ {
		fmt.Printf("This is Outer loop: %d\n", i)
		for a := 1; a <= 5; a++ {
			fmt.Printf("\t\t This is inner loop: %d\n", a)
		}
	}

}
