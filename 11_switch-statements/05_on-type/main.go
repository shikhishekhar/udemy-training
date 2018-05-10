package main

import "fmt"

func main() {

	var x interface{} = 10

	switch x.(type) {
	case int:
		fmt.Println("Welcome Shekhar")
	case string:
		fmt.Println("Welcome Shikhi")
	default:
		fmt.Println("Not a valid option")
	}
}
