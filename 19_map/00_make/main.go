package main

import "fmt"

func main(){
	var greetingMap = make(map[string]string)
	greetingMap["1"] = "Shekhar"
	greetingMap["2"] = "Tikoo"

	fmt.Println(greetingMap)
}