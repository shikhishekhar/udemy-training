package main

import "fmt"

func main() {
	mySlice := make([]int,1, 2)

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}
