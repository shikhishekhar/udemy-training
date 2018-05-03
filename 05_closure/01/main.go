package _1

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "Shikhi Shekhar"
		fmt.Println(y)
	}
	// fmt.println(y) // outside scope of y
}
