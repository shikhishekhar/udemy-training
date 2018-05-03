package main

import "fmt"

func main() {
	a := 43
	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)

	var b *int = &a

	fmt.Println(b)

	fmt.Println(*b)
}
