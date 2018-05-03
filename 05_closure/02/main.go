package main

import "fmt"

func main() {
	x := 42

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
