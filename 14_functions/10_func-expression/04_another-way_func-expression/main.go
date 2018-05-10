package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Shikhi Shekhar"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
}
