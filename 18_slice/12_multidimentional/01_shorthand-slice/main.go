package main

import "fmt"

func main() {
	student := []string{}
	students := [][]string{}

	//student[0] = "Shikhi"
	student = append(student, "Shikhi")

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
