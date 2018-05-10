package main

import "fmt"

func main(){
	numbers := []int{1,2,3,4,5,6}
	fmt.Println(findGreatestNumber(numbers...))
}

func findGreatestNumber(numbers ...int) int{
	highestNumber := 0
	for _, v := range numbers{
		if v > highestNumber{
			highestNumber = v
		}
	}
	return highestNumber
}