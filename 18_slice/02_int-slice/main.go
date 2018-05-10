package main

import "fmt"

func main(){
	mySlice := []string{"a","b","c","d","e"}
	fmt.Println(mySlice)        // [a,b,c,d,e]
	fmt.Println(mySlice[2:4])   // slicing the slice - [c,d]
	fmt.Println(mySlice[2])     // index access; [c]
	fmt.Println("myString"[2])  // index accessing [83]
}