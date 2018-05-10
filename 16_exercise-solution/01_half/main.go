package main

import "fmt"

func main() {
	h, even := divideAndReturnIfDivisibleBy2(2)
	fmt.Println(h, even)
}

func divideAndReturnIfDivisibleBy2(number int) (int, bool) {
	return number/2, number%2 == 0
}
