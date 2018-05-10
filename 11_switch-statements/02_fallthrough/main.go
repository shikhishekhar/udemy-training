package main

import "fmt"

func main() {
	switch "Shikhi" {
	case "Shekhar":
		fmt.Println("Welcome Shekhar")
	case "Shikhi":
		fmt.Println("Welcome Shikhi")
		fallthrough
	case "Tikoo":
		fmt.Println("Welcome Tikoo")
	default:
		fmt.Println("Not a valid option")
	}
}
