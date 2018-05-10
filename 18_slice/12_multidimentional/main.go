package main

import "fmt"

func main() {
	records := make([][]string, 0)

	student1 := make([]string, 1)
	student1[0] = "Shikhi"
	records = append(records, student1)

	student2 := make([]string, 2)
	student2[0] = "Shekhar"
	records = append(records, student2)

	fmt.Println(records)
}
