package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(firstName string, lastName string) (string, string) {
	return fmt.Sprint(firstName, lastName), fmt.Sprint(lastName, firstName)
}
