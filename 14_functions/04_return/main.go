package main

import "fmt"

func main() {
	fmt.Println(greet("Shikhi", "Shekhar"))
}

func greet(firstName string, lastName string) string {
	return fmt.Sprint(firstName, lastName)
}
