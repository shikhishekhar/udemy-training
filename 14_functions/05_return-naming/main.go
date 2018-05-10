package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(firstName string, lastName string) (s string) {
	s = fmt.Sprint(firstName, lastName)
	return
}
