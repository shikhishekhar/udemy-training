package main

import "fmt"

func main() {

	half:= func(number int) (int, bool){
		return number/2, number%2 == 0
	}
	fmt.Println(half(2))
}