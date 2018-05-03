package main

import "fmt"

func main() {
	a := 43
	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 42 // changes the value at this address to 42 instead of 43

	fmt.Println("After value changed using pointer - ", a)
}
