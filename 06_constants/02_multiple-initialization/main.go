package main

import "fmt"

const (
	tempConst string = "Shikhi & Shekhar"
	pi               = 3.14
)

func main() {
	const tempConst2 = 32

	fmt.Println("tempConst2 - ", tempConst2)
	fmt.Println("tempConst - ", tempConst)
	fmt.Println("pi - ", pi)
}
