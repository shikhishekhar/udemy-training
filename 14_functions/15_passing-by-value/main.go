package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age)

	changeAge(&age)

	fmt.Println(age)
	fmt.Println(&age)
}

func changeAge(x *int) {
	fmt.Println(x)
	fmt.Println(*x)

	*x = 24

	fmt.Println(x)
	fmt.Println(*x)
}
