package main

import "fmt"

func main() {
	for i := 1; i <= 2; i++ {
		for j := 11; j <= 12; j++ {
			fmt.Println(i, "-", j)
		}
	}
}
