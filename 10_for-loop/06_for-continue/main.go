package main

import "fmt"

func main() {
	i := 1
	for {
		i++
		if i%2 == 0 {
			continue
		}

		if i%2 != 0 {
			{
				fmt.Println(i)
			}

		}
		if i >= 11 {
			break
		}
	}
}
